/*
 * password.go 
 *
 * History:
 *  1   Jun11 MR Initial version, limited testing
 *  2   Aug13 MR New and much shorter version using bcrypt package
 */

package artistic

import (
	"fmt"
    "code.google.com/p/go.crypto/bcrypt"
)

/*
 * Password - a type defining a user password
 *
 * This struct has only one field. We hide this field into struct, so it 
 * cannot be accessible from outside. We define only operations on password:
 * Get, Set and Cmp (compare) to arbitrary string. From outside we deal with
 * strings, while internally everything is done using bytes.
 * Also, when password is set, it is hashed and stored that way. All operations
 * on passwords are performed using hashed passwords.
 */
type Password struct {
	pwd []byte
}

/*
 * Password.Set - set a password
 *
 * Password will, of course, be hashed using bcrypt algorithm.
 */
func (p *Password) Set(passwd string) (e error) {
    p.pwd, e = bcrypt.GenerateFromPassword([]byte(passwd), 0)
    return e
}

/*
 * Password.Get - return a stored hashed password
 */
func (p *Password) Get() string { return fmt.Sprintf("%x", p.pwd) }

/*
 * Password.Cmp - compare arbitrary password to the one stored 
 */
func (p *Password) Cmp(toCompare string) bool {
    status := true
    e := bcrypt.CompareHashAndPassword(p.pwd, []byte(toCompare))
    if e != nil { status = false }
    return status
}

