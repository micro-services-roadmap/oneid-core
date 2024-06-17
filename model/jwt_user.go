package model

import "github.com/google/uuid"

type JwtUser struct {
	Id          *int64    `json:"id"`
	Uuid        uuid.UUID `json:"uuid"`
	Name        *string   `json:"name"`
	Email       *string   `json:"email"`
	AuthorityId *int      `json:"authorityId"`
}
