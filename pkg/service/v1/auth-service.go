package v1

import (
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/synthesis-backend/auth-service/pkg/api/v1"
	"golang.org/x/net/context"
	"log"
)

var jwtKey = []byte("test_key")

type UserInfo struct {
	Username    string   `json:"user"`
	Groups      []string `json:"groups"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type Claims struct {
	UserInfo
	jwt.StandardClaims
}

type Server struct{}

func (s *Server) CheckAuth(ctx context.Context, in *v1.AuthRequest) (*v1.AuthResponse, error) {

	tokenString := in.Token

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if !token.Valid {
		log.Printf("Invalid token")
		return &v1.AuthResponse{
			Status: v1.AuthResponse_UNAUTHORIZED,
		}, nil
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return &v1.AuthResponse{
				Status: v1.AuthResponse_UNAUTHORIZED,
			}, nil
		}
		return &v1.AuthResponse{
			Status:    v1.AuthResponse_ERROR,
			ErrorCode: v1.AuthResponse_BAD_REQUEST,
		}, nil
	}
	log.Printf("User %s successfully verified. Permissions are: \n %v", claims.Username, claims)
	return &v1.AuthResponse{Status: v1.AuthResponse_APPROVED}, nil
}
