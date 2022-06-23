package server

import (
	"github.com/form3tech-oss/jwt-go"
	"time"
)

func CreateToken(userID string, firstName string) {

	// JOSEヘッダー
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// JWSペイロード

	token.Claims = jwt.MapClaims{
		"iss":  firstName,
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
}
