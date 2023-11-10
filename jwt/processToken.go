package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/nehonar/palominos_algorithm/models"
)

var Email string
var UserId string

func ProcessToken(token string, JWTSign string) (*models.Claim, bool, string, error) {
	personal_key := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	parsed_token, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return personal_key, nil
	})
	if !parsed_token.Valid {
		return &claims, false, string(""), errors.New("Invalid token")
	}

	return &claims, false, string(""), err
}
