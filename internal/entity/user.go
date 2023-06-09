package entity

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint64
	Email    string
	Username string
	Password string
	Bio      sql.NullString
	Image    sql.NullString
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

type Output struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type UserOutput struct {
	Output `json:"user"`
}

func (u *User) PrepareOutput() UserOutput {
	user := Output{
		u.Email,
		u.Username,
		"",
		u.Bio.String,
		u.Image.String,
	}
	return UserOutput{user}
}
