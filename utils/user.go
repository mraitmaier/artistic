package utils
// Package user implements the type that defines a user with username, password, role and usual other information about users.

import (
	"fmt"
	"strings"
)

// AllowedRoles defines a closure that checks which roles are allowed and which are not; used when trying to change the role.
type AllowedRoles func(string) bool

// User defines the user.
// NOTE: we do not allow changing of the username, while role and password change are controlled by special methods.
type User struct {

	// username field is private and cannot be changed
	Username string

	// password field is private so that all operations on password are controlled by the code in this file
    Password string

	// role is a private field so that all changes are controlled by code in this file
	Role string

	// Fullname is user's full name
	Fullname string

	// email is user's email address
	Email string

	// Phone is user's phone number
	Phone string

	// Disabled flag signals that user is disabled and is not allowed to login
	Disabled bool

	// MustChangePassword is a flag that signals user must change his/her password at next login
	MustChangePassword bool

	// AllowedRoles a closure that checks allowed roles for users
    AllowedRoles  `bson:"-"`

    // Visible is a flag - that might be used or not - indicating the user's visibility in the system (UI). 
    // This is actually of limited use: it may be used for a user that cannot be deleted (something like 'god' user).
    // The user is visible by default, of course.
    Visible bool
}

// NewUser creates a new User with only basic (mandatory!) properties defined.
// The 'create' flag is used to indicate whether this is a completely new user and therefore password must be hashed (using
// bcrypt) or it is the existing user with already hashed password (read from config file, for instance). The 'checkrole'
// argument is a closure that will check for allowed roles.
func NewUser(name, passwd, role string, create bool, checkrole AllowedRoles) *User {
	p, _ := NewPassword(passwd, create)
	return &User{name, p.Get(), role, "", "", "", false, false, checkrole, true}
}

// CreateUser creates a new User with all properties defined.
// The 'create' flag is used to indicate whether this is a completely new user and therefore password must be hashed (using
// bcrypt) or it is the already existing user with hashed password (read from config file, for instance). The 'checkrole'
// argument is a closure that will check for allowed roles.
func CreateUser(name, passwd, role, fullname, email, phone string, create bool, checkrole AllowedRoles) *User {
	u := NewUser(name, passwd, role, create, checkrole)
	u.Fullname = fullname
	u.Email = email
	u.Phone = phone
	u.Disabled = false
	u.MustChangePassword = false
    u.Visible = true
	return u
}

// String returns a human-readable representation of the User type.
func (u *User) String() string {
	s := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%t|%t\n", u.Username, u.Password, u.Role,
		u.Fullname, u.Email, u.Phone, u.Disabled, u.MustChangePassword)
	return s
}

// Debug returns a human-readable representation of the User type (used for debugging and testing).
func (u *User) Debug() string {
	s := []string{"User instance:",
		fmt.Sprintf("        username: %s", u.Username),
		fmt.Sprintf("        password: %s", u.Password),
		fmt.Sprintf("            role: %s", u.Role),
		fmt.Sprintf("       full name: %s", u.Fullname),
		fmt.Sprintf("           email: %s", u.Email),
		fmt.Sprintf("           phone: %s", u.Phone),
		fmt.Sprintf("     is disabled: %t", u.Disabled),
		fmt.Sprintf("must change pass: %t", u.MustChangePassword),
		fmt.Sprintf("         visible: %t", u.Visible)}
	return strings.Join(s, "\n")
}

// ChangePassword changes user's password.
// If all input parameters are OK, the user's password is changed.
//
//    INPUT:
//          old - old password (plain string)
//         newp - new password (plain string)
//    repeatedp - new password repeated (plain string)
//
//    RETURNS:
//    error if password has not been changed or nil if all is OK
func (u *User) ChangePassword(old, newp, repeatedp string) error {

	// check if new password strings are identical
	if newp != repeatedp {
		return fmt.Errorf("New password strings do not match.")
	}

    oldpwd, err := NewPassword(u.Password, false)
    if err != nil {
        return err
    }
	// check if we match the old password
	if !oldpwd.Cmp(old) {
		return fmt.Errorf("Invalid old password.")
	}
	// it's OK, change password
	p, err := NewPassword(newp, true)
	if err != nil {
		return err
	}
	u.Password = p.Get()
	return nil
}

// ChangeRole changes the user's role.
// The method first checks the allowed roles closure. If OK, role is changed. Otherwise an error is returned.
func (u *User) ChangeRole(role string) error {

	var e error
	if u.AllowedRoles != nil && u.AllowedRoles(role) {
		u.Role = role
	} else {
		e = fmt.Errorf("Given role is not among allowed roles.")
	}
	return e
}

// ComparePassword does what it says: it compares the given (plain-text) password with existing one.
func (u *User) ComparePassword(toCompare string) bool {

    p, err := NewPassword(u.Password, false)
    if err != nil {
        return false
    }
    return p.Cmp(toCompare)
}

// IsDisabled returns the value of the Disabled flag.
func (u *User) IsDisabled() bool { return u.Disabled }
