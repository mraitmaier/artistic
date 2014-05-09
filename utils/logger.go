/*
 * logger.go -  custom logger implementation
 *
 * Default log module is simply not enough
 *
 * History:
 *  1   Jul11   MR  The initial version
 *  2   May14   MR  Refactoring and simplification: the LogLevel type is out,
 *                  Severity is now used instead. The second is introduction of
 *                  concurency: log can now run as goroutine and messages are
 *                  sent over a channel.
 */

package utils

import (
	"fmt"
	"os"
	"strings"
//    "errors"
	"time"
)

// Converts log level given as string into proper Severity value.
// If invalid string is given, function returns 'UnknownSeverity' value.
func SeverityFromString(lvl string) Severity {
    loglvl := UnknownSeverity
	switch strings.ToUpper(lvl) {
	case "EMERGENCY":
		loglvl = Emergency
	case "ALERT":
		loglvl = Alert
	case "CRITICAL":
		loglvl = Critical
	case "ERROR":
		loglvl = Error
	case "WARNING":
		loglvl = Warning
	case "NOTICE":
		loglvl = Notice
	case "INFO":
		loglvl = Informational
	case "DEBUG":
		loglvl = Debug
	}
	return loglvl
}

/************************** Logger ***********************************/
/*
 * Logger - an interface defining methods for various log handlers
 */
type Logger interface {
	Emergency(string)
	Alert(string)
	Critical(string)
	Error(string)
	Warning(string)
	Notice(string)
	Info(string)
	Debug(string)
	String() string
	Close()
}

/************************** logHandler ***********************************/
/*
 * logHandler - private struct that defines log level and format
 */
type logHandler struct {
	sev  Severity /* set severity for this LogHandler */
	format string   /* a formatter for this LogHandler */
}

func (l *logHandler) Severity() Severity { return l.sev }

func (l *logHandler) SetSeverity(s Severity) { l.sev = s }

func (l *logHandler) Format() string { return l.format }

func (l *logHandler) SetFormat(fmt string) { l.format = fmt }

func newLogHandler(fmt string, sev Severity) *logHandler {
	return &logHandler{sev, fmt}
}

/************************** Log ***********************************/
/*
 * Log - a slice of different Loggers that can be added at will
 */
type logmsg struct {
    sev Severity
    msg string
}

type Log struct {

    // a list of log handlers
	Handlers []Logger

    // a channel to send log messages
    logch chan *logmsg

    // a channel to signal when to stop the logger goroutine
    stop chan int
}

func (l *Log) String() string {
	s := ""
	for _, h := range l.Handlers {
		if h != nil {
			s += fmt.Sprint(h.String())
		}
	}
	return s
}

func (l *Log) findEmpty() int {
	for ix, h := range l.Handlers {
		if h == nil {
			return ix
		}
	}
	return -1
}

func (l *Log) AddHandler(log Logger) []Logger {
	// check length and capacity; resize if needed...
	length := len(l.Handlers)
	c := cap(l.Handlers)
	if length+1 > c {
		newlst := make([]Logger, 0, 2*c)
		copy(newlst, l.Handlers)
		l.Handlers = newlst
	}
	l.Handlers = l.Handlers[0 : length+1]
	// now find empty index and insert Handler
	ix := l.findEmpty()
	if ix > -1 {
		l.Handlers[ix] = log
	}
	return l.Handlers
}

// A dispatch log messages method.
// Calls all needed log handlers and logs the given message with given level.
// If an unknown log level is received, do nothing.
func (l *Log) dispatch(sev Severity, msg string) {
	for _, h := range l.Handlers {
		switch sev {
		case Emergency:
			h.Emergency(msg)
		case Alert:
			h.Alert(msg)
		case Critical:
			h.Critical(msg)
		case Error:
			h.Error(msg)
		case Warning:
			h.Warning(msg)
		case Notice:
			h.Notice(msg)
		case Informational:
			h.Info(msg)
		case Debug:
			h.Debug(msg)
		}
	}
}

func (l *Log) Debug(msg string) { l.send(Debug, msg) }

func (l *Log) Info(msg string) { l.send(Informational, msg) }

func (l *Log) Notice(msg string) { l.send(Notice, msg) }

func (l *Log) Warning(msg string) { l.send(Warning, msg) }

func (l *Log) Error(msg string) { l.send(Error, msg) }

func (l *Log) Critical(msg string) { l.send(Critical, msg) }

func (l *Log) Alert(msg string) { l.send(Alert, msg) }

func (l *Log) Emergency(msg string) { l.send(Emergency, msg) }

func (l *Log) Len() int { return len(l.Handlers) }

func (l *Log) Close() {
	for _, h := range l.Handlers {
		h.Close()
	}

    // send a signal to quit logger goroutine
    if l.stop != nil {
        close(l.logch)
        l.stop <- 1
    }
}

const logLength int = 3

// Create new logger, specify the number of log handlers and create needed  
// channels: the one onto which the log messages are sent and the other where
// signal when to stop is sent.
// Return the Log instance. 
func NewLog(num int) (*Log) {
    // default the handler value
	if num == 0 {
		num = logLength
	}

    // create new Log instance
	l := &Log{ make([]Logger, 0, logLength), nil, nil }
    return l
}

/*
func (l *Log) sendS(sev, msg string) {
    if l.logch != nil {
        l.logch <- &logmsg{ SeverityFromString(sev), msg }
    }
}
*/

func (l *Log) send(sev Severity,  msg string) {
    if l.logch != nil {
        l.logch <- &logmsg{ sev, msg }
    }
}

// run logger as a goroutine
func (l *Log) Run() error {

    // open logger channels 
    l.logch = make(chan *logmsg, 10)  // message channel (buffered)
    l.stop  = make(chan int)          // stop channel

    // now start a logger goroutine
    go func(l *Log) {

        for {
            select {

            // when message is received over channel, write it
            case m :=<-l.logch:
                //fmt.Printf("DEBUG, logger: msg=%v\n", m) // DEBUG
                //l.Log(m.sev, m.msg)
                l.dispatch(m.sev, m.msg)

            // when data is received over stop channel, just exit the goroutine
            case <- l.stop:
                return

            default: // do nothing
            }
        }
    }(l)

    return nil
}

/************************** Formatter  ***********************************/
/*
 * Formatter - an interface defining the generic formatter
 */
type Formatter interface {
	Format(string)
}

/************************** FileHandler ***********************************/
/*
 * FileHandler
 */
type FileHandler struct {
	*logHandler
	file *os.File
}

/*
 * FileHandler.log - creates new stream handler
 */
func (f *FileHandler) log(level Severity, msg string) {
	if f.Severity() >= level {
		fmt.Fprintf(f.file, f.Format(), Now(), level, msg)
	}
}

func (f *FileHandler) Debug(msg string) {
	f.log(Debug, msg)
}

func (f *FileHandler) Info(msg string) {
	f.log(Informational, msg)
}

func (f *FileHandler) Notice(msg string) {
	f.log(Notice, msg)
}

func (f *FileHandler) Warning(msg string) {
	f.log(Warning, msg)
}

func (f *FileHandler) Error(msg string) {
	f.log(Error, msg)
}

func (f *FileHandler) Critical(msg string) {
	f.log(Critical, msg)
}

func (f *FileHandler) Alert(msg string) {
	f.log(Alert, msg)
}

func (f *FileHandler) Emergency(msg string) {
	f.log(Emergency, msg)
}

func (f *FileHandler) Close() {
	if f.file != nil {
		f.file.Close()
	}
}

func (f *FileHandler) String() string {
	return fmt.Sprintf("  FileHandler: fmt=%q, lvl=%-10s, fd=%d\n",
		f.Format(), f.Severity(), f.file.Fd())
}

/*
 * NewFileHandler - creates new file handler
 */
func NewFileHandler(filename string,
	fmt string, sev Severity) (*FileHandler, error) {
	// open log file
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	return &FileHandler{newLogHandler(fmt, sev), f}, err
}

/************************** StreamHandler ***********************************/
/*
 * StreamHandler
 */
type StreamHandler FileHandler

/*
 * StreamHandler.log - creates new stream handler
 */
func (s *StreamHandler) log(level Severity, msg string) {
	if s.Severity() >= level {
		fmt.Printf(s.Format(), Now(), level, msg)
	}
}

func (s *StreamHandler) Debug(msg string) {
	s.log(Debug, msg)
}

func (s *StreamHandler) Info(msg string) {
	s.log(Informational, msg)
}

func (s *StreamHandler) Notice(msg string) {
	s.log(Notice, msg)
}

func (s *StreamHandler) Warning(msg string) {
	s.log(Warning, msg)
}

func (s *StreamHandler) Error(msg string) {
	s.log(Error, msg)
}

func (s *StreamHandler) Critical(msg string) {
	s.log(Critical, msg)
}

func (s *StreamHandler) Alert(msg string) {
	s.log(Alert, msg)
}

func (s *StreamHandler) Emergency(msg string) {
	s.log(Emergency, msg)
}

func (s *StreamHandler) Close() {
	// an empty implementation to satisfy the Logger interface
}

func (s *StreamHandler) String() string {
	return fmt.Sprintf("StreamHandler: fmt=%q, lvl=%-10s\n",
		s.Format(), s.Severity())
}

/*
 * NewStreamHandler - creates new stream handler
 */
func NewStreamHandler(fmt string, sev Severity) *StreamHandler {
	return &StreamHandler{newLogHandler(fmt, sev), os.Stdout}
}

/************************** SyslogHandler ***********************************/
/*
 * SyslogHandler - 
 */
type SyslogHandler struct {
	*logHandler
	IP string
	*SyslogMsg
}

/*
 * SyslogHandler.log - sends a message to the wire using UDP port 514
 */
func (s *SyslogHandler) log(level Severity, msg string) error {
	if s.Severity() >= level {
		s.Fac = FacLocal0
		s.Sev = level
		s.Msg = fmt.Sprintf("%s %s", level.String(), msg)
		t := time.Now()
		s.SetTimestamp(t)
		err := s.Send(s.IP)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (s *SyslogHandler) Debug(msg string) {
	s.log(Debug, msg)
}

func (s *SyslogHandler) Info(msg string) {
	s.log(Informational, msg)
}

func (s *SyslogHandler) Notice(msg string) {
	s.log(Notice, msg)
}

func (s *SyslogHandler) Warning(msg string) {
	s.log(Warning, msg)
}

func (s *SyslogHandler) Error(msg string) {
	s.log(Error, msg)
}

func (s *SyslogHandler) Critical(msg string) {
	s.log(Critical, msg)
}

func (s *SyslogHandler) Alert(msg string) {
	s.log(Alert, msg)
}

func (s *SyslogHandler) Emergency(msg string) {
	s.log(Emergency, msg)
}

func (s *SyslogHandler) Close() {
	// an empty implementation to satisfy the Logger interface
}

func (s *SyslogHandler) String() string {
	return fmt.Sprintf("SyslogHandler: fmt=%q, lvl=%-10s, Server=%q %s\n",
		s.Format(), s.Severity(), s.IP)
}

/*
 * NewSyslogHandler - creates new syslog handler
 */
func NewSyslogHandler(ip string, fmt string, sev Severity) *SyslogHandler {
	//
	return &SyslogHandler{newLogHandler(fmt, sev), ip, NewSyslogMsg()}
}
