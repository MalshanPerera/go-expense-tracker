package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type TokenType struct {
	AccessToken  string
	RefreshToken string
}

var Type = TokenType{
	AccessToken:  "access_token",
	RefreshToken: "refresh_token",
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))
var jwtExpiry time.Duration

func init() {
	expiryHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRED: %v", err)
	}
	jwtExpiry = time.Duration(expiryHours) * time.Hour
}

func CreateToken(userId pgtype.UUID, tokenType string) (string, int64, error) {

	expiredAt := time.Now().Add(jwtExpiry).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":    userId,
			"token_type": tokenType,
			"expired_at": expiredAt,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiredAt, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
