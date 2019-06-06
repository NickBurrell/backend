protoc -I api/proto/v1 \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
       --go_out=plugins=grpc:pkg/api/v1 \
       api/proto/v1/auth-service.proto;

protoc -I api/proto/v1 \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
       --go_out=plugins=grpc:pkg/api/v1 \
       api/proto/v1/health-service.proto;


protoc -I api/proto/v1 \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
       --grpc-gateway_out=logtostderr=true:pkg/api/v1 \
       api/proto/v1/auth-service.proto;


protoc -I api/proto/v1 \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
       --swagger_out=logtostderr=true:api/swagger/v1 \
       api/proto/v1/auth-service.proto;

