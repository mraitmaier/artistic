/*
 * person.go
 */
package artistic

type Artist interface {
    String() string
    Json() (string, error)
}

