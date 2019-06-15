package v1

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/zero-frost/backend/auth-service/pkg/api/v1"
	"github.com/zero-frost/backend/auth-service/pkg/config"
	"github.com/zero-frost/backend/auth-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
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

type AuthServer struct {
	db       *gorm.DB
	settings *config.ServerConfig
}

func NewAuthServer(db *gorm.DB, settings *config.ServerConfig) *AuthServer {
	return &AuthServer{
		db:       db,
		settings: settings,
	}
}

func (s *AuthServer) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	var users []models.User
	tx := s.db.Begin()
	if err := tx.Where("username = ? OR email = ?", in.Username, in.Email).Find(&users).Error; len(users) != 0 {
		for _, elem := range users {
			if elem.Username == in.Username {
				tx.Rollback()
				return &v1.CreateUserResponse{
					Api:       "v1",
					Success:   false,
					ErrorCode: v1.CreateUserResponse_USERNAME_TAKEN,
				}, fmt.Errorf("error: username already taken")
			} else if elem.Email == in.Email {
				tx.Rollback()
				return &v1.CreateUserResponse{
					Api:       "v1",
					Success:   false,
					ErrorCode: v1.CreateUserResponse_EMAIL_IN_USE,
				}, fmt.Errorf("error: email already in use")
			}
		}
	} else if err != nil && err.Error() != "record not found" {
		tx.Rollback()
		return &v1.CreateUserResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.CreateUserResponse_INTERNAL_ERROR,
		}, err
	}

	if !validateEmail(in.Email) {
		tx.Rollback()
		return &v1.CreateUserResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.CreateUserResponse_INVALID_EMAIL,
		}, fmt.Errorf("error: invalid email")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.MinCost)
	if err != nil {
		tx.Rollback()
		return &v1.CreateUserResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.CreateUserResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: failed to hash password")
	}

	if err := tx.Create(&models.User{
		Username: in.Username,
		Email:    in.Email,
		Password: string(passwordHash),
	}).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return &v1.CreateUserResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.CreateUserResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: failed to create new user entry")
	}

	log.Printf("Successfully created user %s", in.Username)
	tx.Commit()
	return &v1.CreateUserResponse{
		Api:     "v1",
		Success: true,
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
	var user models.User
	if in.Username == "" {
		return &v1.LoginResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.LoginResponse_BLANK_USERNAME,
		}, fmt.Errorf("error: no username provided")
	}

	pass := in.Password
	if pass == "" {
		return &v1.LoginResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.LoginResponse_BLANK_PASSWORD,
		}, fmt.Errorf("error: no password provided")
	}
	if err := s.db.Where("username = ?", in.Username).First(&user).Error; err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			{
				return &v1.LoginResponse{
					Api:       "v1",
					Success:   false,
					ErrorCode: v1.LoginResponse_INCORRECT_USERNAME_OR_PASSWORD,
				}, fmt.Errorf("error: incorrect username or password")
			}
		default:
			{
				return &v1.LoginResponse{
					Api:       "v1",
					Success:   false,
					ErrorCode: v1.LoginResponse_BAD_REQUEST,
				}, fmt.Errorf("error: internal server error")
			}
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return &v1.LoginResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.LoginResponse_INCORRECT_USERNAME_OR_PASSWORD,
		}, fmt.Errorf("error: incorrect username or password")
	}

	userData := UserInfo{
		Username: in.Username,
	}

	for _, elem := range user.Roles {
		userData.Roles = append(userData.Roles, elem.RoleID)
	}

	for _, elem := range user.Permissions {
		userData.Permissions = append(userData.Permissions, elem.PermissionID)
	}

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    "Authentication Server",
		},
		UserInfo: userData,
	}

	// var path string
	// {
	// 	usr, err := user.Current()
	// 	if err != nil {
	// 		return &v1.LoginResponse{
	// 			Api:       "v1",
	// 			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
	// 		}, fmt.Errorf("error: could not load secret")
	// 	}
	// 	path = user.HomeDir
	// }
	// file, err := os.Open(path + ".auth-server/secret.json")
	// if err != nil {
	// 	if err != nil {
	// 		return &v1.LoginResponse{
	// 			Api:       "v1",
	// 			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
	// 		}, fmt.Errorf("error: could not load secret")
	// 	}
	// }
	// var secret map[string]interface{}
	// if err := json.NewDecoder(file).Decode(&secret); err != nil {
	// 	return &v1.LoginResponse{
	// 		Api:       "v1",
	// 		ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
	// 	}, fmt.Errorf("error: could not load secret")
	// }
	// if _, ok := secret["jwt_secret"]; !ok {
	// 	return &v1.LoginResponse{
	// 		Api:       "v1",
	// 		ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
	// 	}, fmt.Errorf("error: could not load secret")
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	jwt, err := token.SignedString([]byte("test"))
	if err != nil {
		return &v1.LoginResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: failed to encode JWT")
	}

	if err != nil {
		return &v1.LoginResponse{
			Api:       "v1",
			Success:   false,
			ErrorCode: v1.LoginResponse_INTERNAL_ERROR,
		}, fmt.Errorf("error: failed to generate user token")
	}

	return &v1.LoginResponse{
		Api:     "v1",
		Success: true,
		Token:   jwt,
	}, nil
}

func (s *AuthServer) GetSettings() *config.ServerConfig {
	return s.settings
}
