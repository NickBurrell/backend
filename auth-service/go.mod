module github.com/zero-frost/backend/auth-service

go 1.12

require (
	cloud.google.com/go v0.40.0 // indirect
	github.com/DATA-DOG/go-sqlmock v1.3.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.1
	github.com/jinzhu/gorm v1.9.9
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/zero-frost/backend v0.0.0-20190615030825-366dc0ad1f31
	golang.org/x/crypto v0.1.0
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190611190212-a7e196e89fd3
	google.golang.org/grpc v1.21.1
)

replace github.com/zero-frost/backend => ../
