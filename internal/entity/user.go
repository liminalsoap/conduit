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

func (u *User) PrepareOutput(token string) UserOutput {
	user := Output{
		u.Email,
		u.Username,
		token,
		u.Bio.String,
		u.Image.String,
	}
	return UserOutput{user}
}

type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

type ProfileOutput struct {
	Profile `json:"profile"`
}

func (u *User) PrepareReuseProfileOutput(isFollowing bool) Profile {
	return Profile{
		u.Username,
		u.Bio.String,
		u.Image.String,
		isFollowing,
	}
}

func (u *User) PrepareProfileOutput(isFollowing bool) ProfileOutput {
	profile := u.PrepareReuseProfileOutput(isFollowing)
	return ProfileOutput{profile}
}
