package v1

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"

	"github.com/dgrijalva/jwt-go"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/zero-frost/auth-service/pkg/api/v1"
	"github.com/zero-frost/auth-service/pkg/config"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
}

func Test_AuthService_Create(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening mock database connection", err)
	}
	defer db.Close()

	var gormDB *gorm.DB
	gormDB, err = gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("error: failed to create gorm database connection from mock, '%s'", err)
	}

	s := NewAuthServer(gormDB, &config.ServerConfig{})

	type args struct {
		ctx context.Context
		req *v1.CreateUserRequest
	}
	tests := []struct {
		name    string
		service v1.AuthServer
		args    args
		mock    func()
		want    *v1.CreateUserResponse
		wantErr bool
	}{
		{
			name:    "OK",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.CreateUserRequest{
					Username: "test_user_1",
					Email:    "test_email@gmail.com",
					Password: "test12345",
				},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE (username = $1 OR email = $2)")).WithArgs("test_user_1", "test_email@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"username", "email", "password"}))
				mock.ExpectExec(regexp.QuoteMeta(
					`INSERT INTO "users" ("email","username","password") `+
						`VALUES ($1,$2,$3) `+
						`RETURNING "users".*`)).
					WithArgs("test_email@gmail.com", "test_user_1", sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			want: &v1.CreateUserResponse{
				Api:     "v1",
				Success: true,
			},
			wantErr: false,
		},
		{
			name:    "E-Mail already in use",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.CreateUserRequest{
					Username: "test_user_1",
					Email:    "test_email@gmail.com",
					Password: "test12345",
				},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE (username = $1 OR email = $2)")).WithArgs("test_user_1", "test_email@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"username", "email", "password"}).AddRow("test_user_2", "test_email@gmail.com", ""))
				mock.ExpectRollback()
			},
			want: &v1.CreateUserResponse{
				Api:       "v1",
				ErrorCode: v1.CreateUserResponse_EMAIL_IN_USE,
			},
			wantErr: true,
		},
		{
			name:    "E-Mail already in use",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.CreateUserRequest{
					Username: "test_user_1",
					Email:    "test_email@gmail.com",
					Password: "test12345",
				},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE (username = $1 OR email = $2)")).WithArgs("test_user_1", "test_email@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"username", "email", "password"}).AddRow("test_user_1", "test_email@_2gmail.com", ""))
				mock.ExpectRollback()
			},
			want: &v1.CreateUserResponse{
				Api:       "v1",
				ErrorCode: v1.CreateUserResponse_USERNAME_TAKEN,
			},
			wantErr: true,
		},
		{
			name:    "Invalid E-Mail",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.CreateUserRequest{
					Username: "test_user_1",
					Email:    "test_email",
					Password: "test12345",
				},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE (username = $1 OR email = $2)")).WithArgs("test_user_1", "test_email").WillReturnRows(sqlmock.NewRows([]string{"username", "email", "password"}).AddRow("", "", ""))
				mock.ExpectRollback()
			},
			want: &v1.CreateUserResponse{
				Api:       "v1",
				ErrorCode: v1.CreateUserResponse_INVALID_EMAIL,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.service.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("create user failed '%v'", err)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("create = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AuthService_Login(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening mock database connection", err)
	}
	defer db.Close()

	var gormDB *gorm.DB
	gormDB, err = gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("error: failed to create gorm database connection from mock, '%s'", err)
	}

	s := NewAuthServer(gormDB, &config.ServerConfig{})

	type args struct {
		ctx context.Context
		req *v1.LoginRequest
	}

	tests := []struct {
		name    string
		service v1.AuthServer
		args    args
		mock    func()
		want    *v1.LoginResponse
		wantErr bool
	}{
		{
			name:    "OK",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Username: "test_user_1",
					Password: "test12345",
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE (username = $1)`)).
					WithArgs("test_user_1").
					WillReturnRows(sqlmock.NewRows([]string{"email", "username", "password"}).
						AddRow("test@gmail.com",
							"test_user_1",
							"$2y$12$tgSSbh0evtE.h5V9WbUB2.jX5Nds9GScAdYNBe8CSRfzpgsF2qRri"),
					)
			},
			want: &v1.LoginResponse{
				Api:   "v1",
				Token: "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidGVzdF91c2VyXzEiLCJyb2xlcyI6bnVsbCwicGVybWlzc2lvbnMiOm51bGwsImV4cCI6MTU1OTk1NjIwNSwiaWF0IjoxNTU5OTUyNjA1LCJpc3MiOiJBdXRoZW50aWNhdGlvbiBTZXJ2ZXIifQ.EQpoBdr2GVG9rzevBoLaCOUMNunzRA1TpaIttg0ia8Uz3FtPdBZF2-6XPQXYEF9P",
			},
			wantErr: false,
		},
		{
			name:    "Blank Username",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Username: "",
					Password: "test12345",
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE (username = $1)`)).
					WithArgs("").
					WillReturnRows(sqlmock.NewRows([]string{"email", "username", "password"}).
						AddRow("test@gmail.com",
							"test_user_1",
							"$2y$12$tgSSbh0evtE.h5V9WbUB2.jX5Nds9GScAdYNBe8CSRfzpgsF2qRri"),
					)
			},
			want: &v1.LoginResponse{
				Api:       "v1",
				ErrorCode: v1.LoginResponse_BLANK_USERNAME,
			},
			wantErr: true,
		},
		{
			name:    "Blank Password",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Username: "test_user_1",
					Password: "",
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE (username = $1)`)).
					WithArgs("test_user_1").
					WillReturnRows(sqlmock.NewRows([]string{"email", "username", "password"}).
						AddRow("test@gmail.com",
							"test_user_1",
							"$2y$12$tgSSbh0evtE.h5V9WbUB2.jX5Nds9GScAdYNBe8CSRfzpgsF2qRri"),
					)
			},
			want: &v1.LoginResponse{
				Api:       "v1",
				ErrorCode: v1.LoginResponse_BLANK_USERNAME,
			},
			wantErr: true,
		},
		{
			name:    "Incorrect Username",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Username: "test_user_2",
					Password: "test12345"},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE (username = $1)`)).
					WithArgs("test_user_2").
					WillReturnRows(sqlmock.NewRows([]string{"email", "username", "password"}).
						AddRow("test@gmail.com",
							"test_user_1",
							"$2y$12$tgSSbh0evtE.h5V9WbUB2.jX5Nds9GScAdYNBe8CSRfzpgsF2qRri"),
					)
			},
			want: &v1.LoginResponse{
				Api:       "v1",
				ErrorCode: v1.LoginResponse_INCORRECT_USERNAME_OR_PASSWORD,
			},
			wantErr: true,
		},
		{
			name:    "Incorrect Password",
			service: s,
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Username: "test_user_1",
					Password: "test12346"},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE (username = $1)`)).
					WithArgs("test_user_1").
					WillReturnRows(sqlmock.NewRows([]string{"email", "username", "password"}).
						AddRow("test@gmail.com",
							"test_user_1",
							"$2y$12$tgSSbh0evtE.h5V9WbUB2.jX5Nds9GScAdYNBe8CSRfzpgsF2qRri"),
					)
			},
			want: &v1.LoginResponse{
				Api:       "v1",
				ErrorCode: v1.LoginResponse_INCORRECT_USERNAME_OR_PASSWORD,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.service.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("failed to login '%v'", err)
				return
			}
			if err == nil && got.Token != "" {
				token, jwtErr := jwt.Parse(got.Token, func(token *jwt.Token) (interface{}, error) {
					return []byte("test_key"), nil
				})
				switch jwtErr.(type) {
				case nil:
					var expectedToken *jwt.Token
					expectedToken, jwtErr = jwt.Parse(got.Token, func(token *jwt.Token) (interface{}, error) {
						return []byte("test_key"), nil
					})
					switch jwtErr.(type) {
					case nil:
						if !reflect.DeepEqual(token, expectedToken) {
							t.Errorf("failed to verify equilivent JWT payloads\n")
						}
					case *jwt.ValidationError:
						t.Errorf("failed to expected validate JWT, '%v'\n", err)
					default:
						t.Errorf("failed to expected parse JWT, '%v'\n", err)
					}
				case *jwt.ValidationError:
					t.Errorf("failed to validate JWT, '%v'\n", err)
				default:
					t.Errorf("failed to parse JWT, '%v'\n", err)
				}

			} else if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("create = %v, want %v", got, tt.want)
			}
			t.Logf("%v", got)
		})
	}
}
