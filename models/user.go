package models

import (
	"errors"
	"reflect"
	"strings"
)

const (
	// USERTABLE ...
	USERTABLE = "app_user"
	// INSERT ...
	INSERT = "INSERT INTO "
	// VALUES ...
	VALUES = "VALUES"
	// RETURNING ...
	RETURNING = "RETURNING"
)

var (
	errEmailIsEmpty      = errors.New("email field is empty, should be filled")
	errPasswordIsEmpty   = errors.New("password field is empty, should be filled")
	errEmailContainsChar = errors.New("email doesn't contains @, please fill email field properly")
	errPasswordLength    = errors.New("password is less than 6 characters")
	errPasswordMatch     = errors.New("passwords doesn't match")
)

//User from router
type User struct {
	Email         string `json:"email" db:"email"`
	Password      string `json:"password" db:"password"`
	PasswordCheck string `json:"passwordCheck" db:"-"`
	DisplayName   string `json:"displayName" db:"display_name"`
}

//ValidateFields All User fields
func (u *User) ValidateFields() error {
	if err := u.validateEmail(); err != nil {
		return err
	}
	if err := u.validatePassword(); err != nil {
		return err
	}
	u.validateDisplayName()

	return nil
}

func (u *User) validateEmail() error {
	if u.Email == "" {
		return errEmailIsEmpty
	}
	if !strings.Contains(u.Email, "@") {
		return errEmailContainsChar
	}

	return nil
}

func (u *User) validatePassword() error {
	if u.Password == "" {
		return errPasswordIsEmpty
	}
	if len(u.Password) < 6 {
		return errPasswordLength
	}
	if u.Password != u.PasswordCheck {
		return errPasswordMatch
	}
	return nil
}

func (u *User) validateDisplayName() {
	if u.DisplayName == "" {
		u.DisplayName = u.Email
	}
}

// err := tx.QueryRowx(`INSERT INTO app_user (email, password, display_name) VALUES ($1, $2, $3) RETURNING id, email, password, display_name`, newUser.Email, newUser.Password, newUser.DisplayName

// InsetSQLString ...
func (u *User) InsetSQLString() string {

	sqlText := INSERT + USERTABLE
	sqlTags := ""
	sqlValues := ""

	val := reflect.ValueOf(u).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		if tag.Get("db") != "-" {
			if val.NumField()-i == 1 {
				sqlTags += tag.Get("db")
				sqlValues += "'" + valueField.String() + "'"
			} else {
				sqlTags += tag.Get("db") + ","
				sqlValues += "'" + valueField.String() + "'" + ","
			}
		}

		// fmt.Printf("%s,\t %v,\t %s\n", typeField.Name, valueField.Interface(), tag.Get("json"))
	}
	sqlText += "(" + sqlTags + ") " + VALUES + " (" + sqlValues + ")" + RETURNING + " id, " + sqlTags
	return sqlText
}

//UserFromDB from router
type UserFromDB struct {
	ID          uint64 `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	DisplayName string `json:"displayName" db:"display_name"`
}
