package v1

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/rcrowley/go-metrics"

	"github.com/zero-frost/backend/auth-service/pkg/api/v1"
)

type MetricServer struct{}

func NewMetricServer() *MetricServer {
	return &MetricServer{}
}

func (s *MetricServer) GetMetrics(ctx context.Context, in *v1.MetricsRequest) (*v1.MetricsResponse, error) {

	data := metrics.DefaultRegistry.GetAll()
	var res []*v1.MetricsResponse_Metric
	for name, metric := range data {
		metricType := reflect.TypeOf(
			metrics.DefaultRegistry.Get(name)).String()[9:]
		metricData, err := json.Marshal(metric)
		if err != nil {
			return &v1.MetricsResponse{
				ErrorCode: v1.MetricsResponse_ENCODING_ERROR,
			}, err
		}
		res = append(res, &v1.MetricsResponse_Metric{
			Name:  name,
			Type:  metricType,
			Value: string(metricData),
		})
	}
	return &v1.MetricsResponse{
		Metrics: res,
	}, nil
}
