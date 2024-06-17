package model

import (
	"encoding/json"
	"github.com/google/uuid"
)

type JwtUser struct {
	Id          *int64    `json:"id"`
	Uuid        uuid.UUID `json:"uuid"`
	Name        *string   `json:"name"`
	Email       *string   `json:"email"`
	AuthorityId *int      `json:"authorityId"`
}

func (u *JwtUser) UserMarshal() string {
	marshal, err := json.Marshal(u)
	if err != nil {
		panic("Marshal JwtValueOfUser failed: " + err.Error())
	}

	return string(marshal)
}

func UserUnMarshal(value string) *JwtUser {
	var u JwtUser
	err := json.Unmarshal([]byte(value), &u)
	if err != nil {
		panic("Unmarshal claims.Value failed: " + err.Error())
	}

	return &u
}
