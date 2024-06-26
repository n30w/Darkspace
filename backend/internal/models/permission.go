package models

import (
	"errors"
	"strings"
)

// member defines the affiliation of a user, whether they are a student,
// a teacher, or an administrator. In other words,
// it defines what group someone is a part of.
// The frontend will send either a 0 for STUDENT or a 1 for TEACHER.
// The affiliation ADMIN is only created server-side.
// The member type implements the Credential interface.
type member uint8

const (
	STUDENT member = iota
	TEACHER
	ADMIN
)

// String returns the string representation of the membership type.
func (m member) String() string {
	switch m {
	case STUDENT:
		return "STUDENT"
	case TEACHER:
		return "TEACHER"
	case ADMIN:
		return "ADMIN"
	default:
		return ""
	}
}

// Valid returns an error, checking whether the membership value
// provided is even valid.
func (m member) Valid() error {
	if m > 2 || m < 0 {
		return errors.New("invalid membership enumeration")
	}

	return nil
}

type scope uint8

const (
	// Determines what one can do with themselves.

	SELF scope = iota

	// Scopes for general pedagogy.

	COURSE
	QUIZ
	ASSIGNMENT
	DISCUSSION
	PROJECT

	// Object specific and contextual scopes.

	COMMENT
	MEDIAS // MEDIAS has an "s", in order to differentiate between MEDIA
	SUBMIT
)

// Permissions dictate what a user can do. Permissions can be overridden.
// Permissions do not only need to exist on the user struct but also as
// an attribute to assignments or other pedagogical structures or discussions.
// If permissions exist on such a structure, they take precedence over
// user permissions.

// permissions is a map of scopes that are keys for permission
// values.
type permissions map[scope]permission

// var permissionsForUser permissions = map[scope]{}
// permissionsForUser[COMMENT] <- accesses comment permissions
// permissionsForUser[COURSE].read <- accesses read field for course permissions

func createPermissions() permissions {
	return permissions{
		SELF:       permission{},
		COURSE:     permission{},
		QUIZ:       permission{},
		ASSIGNMENT: permission{},
		DISCUSSION: permission{},
		PROJECT:    permission{},
		COMMENT:    permission{},
		MEDIAS:     permission{},
		SUBMIT:     permission{},
	}
}

// permission is a permission.
// Remember that Go has zero values that equate to
// false for boolean types. Therefore, default permissions
// means that an object has no permission to do any type of data
// manipulation, creation, or reading.
type permission struct {
	read, write, update, delete bool
}

func newPermission(r, w, u, d bool) permission {
	return permission{
		read:   r,
		write:  w,
		update: u,
		delete: d,
	}
}

// Permissions are stored in the database as a string.

// String serializes the permissions to store in a textual database field.
func (p permission) String() string {
	s := []string{"-", "-", "-", "-"}

	if p.read {
		s[0] = "r"
	}

	if p.write {
		s[1] = "w"
	}

	if p.update {
		s[2] = "u"
	}

	if p.delete {
		s[3] = "d"
	}

	return strings.Join(s, "")
}

// fromString creates a new permission object from a
// string. The string could really just be turned into a hash
// from JSON data but maybe. Who knows.
func fromString(s string) permission {

	p := permission{}

	dashCheck := func(s byte) bool {
		return s != '-'
	}

	if dashCheck(s[0]) {
		p.read = true
	}

	if dashCheck(s[1]) {
		p.write = true
	}

	if dashCheck(s[2]) {
		p.update = true
	}

	if dashCheck(s[3]) {
		p.delete = true
	}

	return p
}

// AccessControl defines a user's membership and their permissions.
// Essentially, the scope of their abilities. It abstracts away type permissions
// in order to conceal behavior and interference in the higher levels
// of the API. AccessControl exists on pedagogical types or users as a pointer.
// Reason being, remember that in Go, pointers have a null value as their
// zero value. This means that a null value has the meaning that an object
// has no permissions at all.
type AccessControl struct {
	perms permissions
}

func (a AccessControl) Read(s scope) bool {
	return a.perms[s].read
}

func (a AccessControl) Write(s scope) bool {
	return a.perms[s].write
}

func (a AccessControl) Update(s scope) bool {
	return a.perms[s].update
}

func (a AccessControl) Delete(s scope) bool {
	return a.perms[s].delete
}
