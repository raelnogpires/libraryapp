package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	secret string
	issure string
}

var str string = os.Getenv("JWT_SECRET")

func NewJWTService() *JWTService {
	return &JWTService{
		secret: str,
		issure: "library",
	}
}

type Claim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

func (j *JWTService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
			Issuer:    j.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (j *JWTService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// https://pkg.go.dev/github.com/golang-jwt/jwt#Parse
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}

		return []byte(j.secret), nil
	})

	return err == nil
}

func (j *JWTService) GetIdFromToken(tkn string) (int64, error) {
	var hmacSampleSecret []byte

	token, err := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}

		return hmacSampleSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["sub"].(string)
		val, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return 0, err
		}

		return val, err
	}

	return 0, err
}
