package login

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/neandrson/go-daev4/orchestrator/entities"
	"github.com/neandrson/go-daev4/orchestrator/utils"
)

type (
	Service struct {
		user UserLogin
	}
)

func NewService(user UserLogin) *Service {
	return &Service{
		user: user,
	}
}

func (s *Service) Login(user entities.User) *Response {
	_, err := s.user.Login(&user)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	hmacSampleSecret := utils.Secret
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"nbf":   now.Unix(),
		"exp":   now.Add(5 * time.Hour).Unix(),
		"iat":   now.Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	return &Response{
		Token: tokenString,
		Error: "",
	}
}
