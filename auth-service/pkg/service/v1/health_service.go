package v1

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/zero-frost/backend/auth-service/pkg/api/v1"
)

type HealthServer struct {
	mu        sync.Mutex
	statusMap map[string]v1.HealthCheckResponse_ServingStatus
}

func NewHealthServer() *HealthServer {
	return &HealthServer{
		statusMap: make(map[string]v1.HealthCheckResponse_ServingStatus),
	}
}

func (s *HealthServer) Check(ctx context.Context, in *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if in.Service == "" {
		return &v1.HealthCheckResponse{
			Status: v1.HealthCheckResponse_SERVING,
		}, nil
	}
	if status, ok := s.statusMap[in.Service]; ok {
		return &v1.HealthCheckResponse{
			Status: status,
		}, nil
	}
	return nil, grpc.Errorf(codes.NotFound, "unknownservice")
}

func (s *HealthServer) SetServingStatus(service string, status v1.HealthCheckResponse_ServingStatus) {
	s.mu.Lock()
	s.statusMap[service] = status
	s.mu.Unlock()
}

func (s *HealthServer) Watch(in *v1.HealthCheckRequest, stream v1.Health_WatchServer) error {
	for _, elem := range s.statusMap {
		if err := stream.Send(&v1.HealthCheckResponse{Status: elem}); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}
