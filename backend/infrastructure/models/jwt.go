package models

import "github.com/google/uuid"

type GetDataToken struct {
	Account_id uuid.UUID `json:"account_id,omitempty"`
}
