package domain

import (
	"fmt"
	"reflect"

	"github.com/n30w/Darkspace/internal/models"
)

type UserStore interface {
	InsertUser(u *models.User) error
	GetUserByID(userid string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	DeleteCourseFromUser(u *models.User, courseid models.CourseId) error
}

type UserService struct {
	store UserStore
}

func NewUserService(us UserStore) *UserService {
	return &UserService{store: us}
}

// CreateUser validates User model values, and if all is well,
// creates the user in the database.
func (us *UserService) CreateUser(um *models.User) error {
	// First check if user exists.
	_, err := us.store.GetUserByID(um.ID)
	if err != nil {
		return err
	}

	// Check if credentials are valid.
	err = validateCredentials(um)
	if err != nil {
		return err
	}

	// Check if email is already in use.
	_, err = us.store.GetUserByEmail(um.Email.String())
	if err == nil {
		return err
	}

	// Check if username is already in use.
	_, err = us.store.GetUserByUsername(um.Username.String())
	// Notice that err IS EQUAL TO nil and not NOT EQUAL TO.
	if err == nil {
		return err
	}

	// If all is well...
	err = us.store.InsertUser(um)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetByID(userid string) (*models.User, error) {
	user, err := us.store.GetUserByID(userid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// What if we want only some information from Assignments or Courses?
func (us *UserService) RetrieveFromUser(
	userid string,
	field string,
) (interface{}, error) {
	user, err := us.store.GetUserByID(userid)
	if err != nil {
		return nil, err
	}

	model := reflect.ValueOf(user)
	fieldValue := model.FieldByName(field)

	if fieldValue == reflect.ValueOf(nil) {
		return nil, fmt.Errorf(
			"field %s does not exist or is uninitialized",
			field,
		)
	}
	return fieldValue, nil

}

func (us *UserService) UnenrollUserFromCourse(
	userid string,
	courseid models.CourseId,
) error {
	user, err := us.store.GetUserByID(userid)
	if err != nil {
		return err
	}
	err = us.store.DeleteCourseFromUser(user, courseid)
	if err != nil {
		return err
	}
	return nil

}

func (us *UserService) NewUsername(s string) Username {
	return Username(s)
}

func (us *UserService) NewPassword(s string) Password {
	return Password(s)
}

func (us *UserService) NewEmail(s string) Email {
	return Email(s)
}

func (us *UserService) NewMembership(d int) Membership {
	return Membership(d)
}
