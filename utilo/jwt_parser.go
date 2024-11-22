package utilo

import (
	"crypto/rsa"
	"errors"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWTParser struct {
	SigningKey *rsa.PublicKey
}

func GetParserOfAbPath(abPath string) (*JWTParser, error) {

	if pk, err := ReadPubKeyOfAbPath(abPath); err != nil {
		return nil, errors.New("读取公钥失败: " + err.Error())
	} else {
		return &JWTParser{pk}, nil
	}
}

// Deprecated: GetParserOfAbPath(相对路径)
func GetParser(path string) (*JWTParser, error) {

	if pk, err := ReadPubKey(path); err != nil {
		return nil, errors.New("读取公钥失败: " + err.Error())
	} else {
		return &JWTParser{pk}, nil
	}
}

func GetParserFromKey(pubKey string) (*JWTParser, error) {

	if pk, err := BuildPubKey(pubKey); err != nil {
		return nil, errors.New("读取公钥失败: " + err.Error())
	} else {
		return &JWTParser{pk}, nil
	}
}

// ParseToken 解析 token
func (j *JWTParser) ParseToken(tokenString string) (*modelx.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &modelx.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet // can do not validate jwt start time
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*modelx.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
