package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Account_id     uuid.UUID `json:"account_id,omitempty"`
	Account_create time.Time `json:"time,omitempty"`
	Email          string    `json:"email" validate:"required,email"`
	Username       string    `json:"username" validate:"required"`
	Password       string    `json:"password" validate:"required"`
	Birth_date     string    `json:"birth_date,omitempty"`
	// Profile        Profile
}

type Login struct {
	Email    string `json:"email,omitempty" validate:"email"`
	Username string `json:"username,omitempty"`
	Password string `json:"password" validate:"required"`
}

type PublicUser struct {
	Account_create time.Time `json:"time,omitempty"`
	Username       string    `json:"username" validate:"required"`
}

// type Profile struct {
// 	Username string
// 	Gender   string
// 	Quote    string
// }

type Contents struct {
	Content_id     uuid.UUID `json:"content_id,omitempty"`
	Account_create time.Time `json:"time,omitempty"`
	Account_id     uuid.UUID `json:"account_id,omitempty"`
	Content_type   string    `json:"content_type" validate:"required"`
	Title          string    `json:"title" validate:"required"`
	Contents       string    `json:"contents" validate:"required"`
}
