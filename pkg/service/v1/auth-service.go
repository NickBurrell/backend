package v1

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"gitlab.com/zero_frost/auth-service/models"
	"gitlab.com/zero_frost/auth-service/pkg/api/v1"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/user"
	"regexp"
	"time"
)

func validateEmail(email string) bool {

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)

}

var jwtKey = []byte("test_key")

type UserInfo struct {
	Username    string `json:"user"`
	Roles       []int  `json:"roles"`
	Permissions []int  `json:"permissions"`
}

type Claims struct {
	UserInfo
	jwt.StandardClaims
}

type Server struct {
	db    *models.Datastore
	cache *redis.Client
}

func (s *Server) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
	user := models.User{
		Username: in.Username,
	}
	if err := s.db.GetUser(&user); err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			{
				if in.Username == "" {
					return &v1.LoginResponse{
						Api:       "v1",
						ErrorCode: v1.LoginResponse_BLANK_USERNAME,
					}, err
				} else {
					return &v1.LoginResponse{
						Api:       "v1",
						ErrorCode: v1.LoginResponse_INCORRECT_USERNAME_OR_PASSWORD,
					}, err
				}
			}
		default:
			{
				return &v1.LoginResponse{
					Api:       "v1",
					ErrorCode: v1.LoginResponse_BAD_REQUEST,
				}, err
			}
		}
	}
	pass := in.Password
	if pass == "" {
		return &v1.LoginResponse{
			Api:       "v1",
			ErrorCode: v1.LoginResponse_BLANK_PASSWORD,
		}, fmt.Errorf("error: no password provided")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return &v1.LoginResponse{
			Api:       "v1",
			ErrorCode: v1.LoginResponse_INCORRECT_USERNAME_OR_PASSWORD,
		}, fmt.Errorf("error: incorrect password")
	}

	userData := UserInfo{
		Username: in.Username,
	}

	for _, elem := range user.Roles {
		userData.Roles = append(userData.Roles, elem.ResourceID)
	}

	for _, elem := range user.Permissions {
		userData.Permissions = append(userData.Permissions, elem.ResourceID)
	}

	claims := Claims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
			Issuer:    "Authentication Server",
		},
		userInfo,
	}

	var path string
	{
		usr, err := user.Current()
		if err != nil {
			return &v1.LoginResponse{
				Api:       "v1",
				ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
			}, fmt.Errorf("error: could not load secret")
		}
		path = user.HomeDir
	}
	file, err := os.Open(path + ".auth-server/secret.json")
	if err != nil {
		if err != nil {
			return &v1.LoginResponse{
				Api:       "v1",
				ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
			}, fmt.Errorf("error: could not load secret")
		}
	}
	var secret map[string]interface{}
	if err := json.NewDecoder(file).Decode(&secret); err != nil {
		return &v1.LoginResponse{
			Api:       "v1",
			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: could not load secret")
	}
	if ok, _ := secret["jwt_secret"]; !ok {
		return &v1.LoginResponse{
			Api:       "v1",
			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: could not load secret")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(secret[jwt_secret])
	if err != nil {
		return &v1.LoginResponse{
			Api:       "v1",
			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: failed to encode JWT")
	}

	opaqueToken, err := uuid.NewV4()
	if err != nil {
		return &v1.LoginResponse{
			Api:       "v1",
			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: failed to generate user token")
	}

	s.cache.Set(string(opaqueToken), tokenstring, claims.ExpiresAt)

	return &v1.LoginResponse{
		Api:         "v1",
		opaqueToken: opaqueToken,
	}

}

func (s *Server) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
}
