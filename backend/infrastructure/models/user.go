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
	// Birth_date     string    `json:"birth_date,omitempty"`
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
	Content_create time.Time `json:"content_create,omitempty"`
	Content_edit   time.Time `json:"content_edit,omitempty"`
	Account_id     uuid.UUID `json:"account_id,omitempty"`
	Content_type   string    `json:"content_type" validate:"required"`
	Title          string    `json:"title" validate:"required"`
	Content_data   string    `json:"contents" validate:"required"`
}

type Userconten struct {
	Username       string    `json:"username,omitempty"`
	Content_id     uuid.UUID `json:"content_id,omitempty"`
	Content_create time.Time `json:"content_create,omitempty"`
	Content_type   string    `json:"content_type" validate:"required"`
	Title          string    `json:"title" validate:"required"`
	Content_data   string    `json:"contents" validate:"required"`
}

type GetContent struct {
	Content_id   uuid.UUID `json:"content_id,omitempty"`
	Content_type string    `json:"content_type,omitempty"`
	Title        string    `json:"title,omitempty"`
	Content_data string    `json:"contents,omitempty"`
}

type Profile struct {
	Profile_id   uuid.UUID
	Account_id   uuid.UUID
	Birth_date   time.Time
	Contact_info int
	Locations    string
	Education    string
	About        string
}
