package model

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

const (
	ID          = "ID"
	UID         = "UID"
	Name        = "NAME"
	Email       = "EMAIL"
	AuthorityId = "AUTHORITY_ID"
)

type JwtUser struct {
	Id          int64     `json:"id"`
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	AuthorityId int64     `json:"authorityId"`
}

func (u *JwtUser) UserMarshal() string {
	marshal, err := json.Marshal(u)
	if err != nil {
		panic("Marshal JwtValueOfUser failed: " + err.Error())
	}

	return string(marshal)
}

func MustUser(value string) *JwtUser {

	if u, err := UserUnMarshal(value); err != nil {
		panic(err.Error())
	} else {
		return u
	}
}

func UserUnMarshal(value string) (*JwtUser, error) {
	var u JwtUser
	err := json.Unmarshal([]byte(value), &u)
	if err != nil {
		return nil, errors.New("Unmarshal claims.Value failed: " + err.Error())
	}

	return &u, nil
}

func GetUserID(value string) (int64, error) {
	if u, err := UserUnMarshal(value); err != nil {
		return 0, err
	} else {
		return u.Id, nil
	}
}

func TryUserID(value string, defaultVal ...int64) int64 {
	if u, err := UserUnMarshal(value); err != nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		} else {
			return -1
		}
	} else {
		return u.Id
	}
}

func MustUserID(value string) int64 {
	return MustUser(value).Id
}

func GetUID(value string) (uuid.UUID, error) {
	if u, err := UserUnMarshal(value); err != nil {
		return uuid.Nil, err
	} else {
		return u.Uuid, nil
	}
}

func TryUID(value string) uuid.UUID {
	if u, err := UserUnMarshal(value); err != nil {
		return uuid.Nil
	} else {
		return u.Uuid
	}
}

func MustUID(value string) uuid.UUID {
	return MustUser(value).Uuid
}

func GetAuthorityId(value string) (int64, error) {
	if u, err := UserUnMarshal(value); err != nil {
		return 0, err
	} else {
		return u.AuthorityId, nil
	}
}

func TryAuthorityId(value string, defaultVal ...int64) int64 {
	if u, err := UserUnMarshal(value); err != nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		} else {
			return -1
		}
	} else {
		return u.AuthorityId
	}
}

func MustAuthorityId(value string) int64 {
	return MustUser(value).AuthorityId
}

func GetEmail(value string) (string, error) {
	if u, err := UserUnMarshal(value); err != nil {
		return "", err
	} else {
		return u.Email, nil
	}
}

func TryEmail(value string, defaultVal ...string) string {
	if u, err := UserUnMarshal(value); err != nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		} else {
			return ""
		}
	} else {
		return u.Email
	}
}

func MustEmail(value string) string {
	return MustUser(value).Email
}

func GetName(value string) (string, error) {
	if u, err := UserUnMarshal(value); err != nil {
		return "", err
	} else {
		return u.Name, nil
	}
}

func TryName(value string, defaultVal ...string) string {
	if u, err := UserUnMarshal(value); err != nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		} else {
			return ""
		}
	} else {
		return u.Name
	}
}

func MustName(value string) string {
	return MustUser(value).Name
}
