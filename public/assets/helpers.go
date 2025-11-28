package auth_gateway

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ValidateAuthRequest(r *AuthRequest) error {
	return validator.New().Struct(r)
}

func GenerateToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user,
	})
	return token.SignedString([]byte("secret"))
}

func DecodeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	return token, err
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", ErrNoToken
	}
	return tokenString, nil
}

func GetUsernameFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if!ok ||!token.Valid {
		return "", ErrInvalidToken
	}
	return claims["username"].(string), nil
}

type Err struct {
	Code    int
	Message string
}

func (e *Err) Error() string {
	return e.Message
}

var (
	ErrNoToken    = &Err{Code: http.StatusUnauthorized, Message: "no token provided"}
	ErrInvalidToken = &Err{Code: http.StatusUnauthorized, Message: "invalid token"}
)