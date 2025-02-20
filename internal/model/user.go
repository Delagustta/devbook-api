package model

import (
	"errors"
	"strings"
	"time"
)

// User represents an user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}
	u.formatFields()
	return nil
}

func (u *User) validate() error {

	if u.Name == "" {
		return errors.New("name is required and cannot be empty")
	}

	if u.Nick == "" {
		return errors.New("nick is required and cannot be empty")
	}

	if u.Email == "" {
		return errors.New("email is required and cannot be empty")
	}

	if u.Password == "" {
		return errors.New("password is required and cannot be empty")
	}

	return nil
}

func (u *User) formatFields() {
	u.Name = strings.Join(strings.Fields(u.Name), " ")
	u.Nick = strings.Join(strings.Fields(u.Nick), " ")
	u.Email = strings.Join(strings.Fields(u.Email), " ")
}
