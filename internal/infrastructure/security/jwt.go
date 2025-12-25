package security

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(userid uint) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load variables from .env file, relying on system envs")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"user_id": userid,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwt_secret))
}

func VerifyToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
