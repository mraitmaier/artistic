package utils
/*
 * password.go 
 *
 * History:
 *  1   Jun11 MR Initial version, limited testing
 *  2   Aug13 MR New and much shorter version using bcrypt package
 *  3   Dec14 MR URL for bcrypt module has changed
 */

import (
	"fmt"
    "golang.org/x/crypto/bcrypt"
)

// Password defines a user password.
// This struct has only one field. We hide this field into struct, so it cannot be accessible from outside. We define only 
// operations on password: Get and Cmp (compare) to arbitrary string. From outside we deal with strings, while internally 
// everything is done using bytes. Also, when password is set, it is hashed and stored that way. Further operations on password 
// are performed using hashes.
type Password struct {
	pwd []byte
}

// NewPassword creates a new password. 
// If the 'toHash' flag is set, the password will, of course, be hashed using the bcrypt algorithm. If not, we assume this is 
// an existing (already hashed) password and the hash is simsimpply stored. 
func NewPassword(passwd string, toHash bool) (pwd *Password, e error) {

    pwd = new(Password)

    if toHash {
        pwd.pwd, e = bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
    } else {
        pwd.pwd = []byte(passwd)
    }
    return
}

// Get returns a stored hashed password.
//func (p *Password) Get() string { return fmt.Sprintf("%x", p.pwd) }
func (p *Password) Get() string { return fmt.Sprintf("%s", p.pwd) }

// Cmp compares arbitrary (clear-text) password to the one stored 
func (p *Password) Cmp(toCompare string) bool {

    if err := bcrypt.CompareHashAndPassword(p.pwd, []byte(toCompare)); err != nil {
        return false
    }
    return true
}

