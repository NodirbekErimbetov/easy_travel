package models

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginUser struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
}
