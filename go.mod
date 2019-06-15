module github.com/zero-frost/backend

go 1.12

require (
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/golang/mock v1.3.1 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/zero-frost/backend/auth-service v0.0.0-20190615030825-366dc0ad1f31 // indirect
	google.golang.org/grpc v1.21.1
)

replace github.com/zero-frost/backend/auth-service => ./auth-service

replace github.com/zero-frost/backend/matchmaking_service => ./matchmaking-service
