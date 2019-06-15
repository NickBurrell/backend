package metrics

import (
	"context"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/rcrowley/go-metrics"
	"google.golang.org/grpc"
)

func splitMethod(s string) string {
	splitMethodName := strings.Split(s, "/")
	splitPreamble := strings.Split(splitMethodName[1], ".")
	return splitPreamble[1] + "." + strings.ToLower(splitPreamble[2]) + "." + strcase.ToSnake(splitMethodName[2])
}

// TODO: Implement persistent logs
func MetricsUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		c := metrics.GetOrRegisterCounter("api."+splitMethod(info.FullMethod)+".count", nil)
		c.Inc(1)

		resp, err := handler(ctx, req)

		return resp, err

	}
}
