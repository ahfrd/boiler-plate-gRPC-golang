package service

import (
	"context"
	"grpc-boiler-plate-go/app/model/proto/health"
	"grpc-boiler-plate-go/app/repository"
)

type (
	healthCheckService struct {
		HealthCheckRepo repository.HealthCheckRepoIn
	}

	HealthCheckServiceIn interface {
		Check(ctx context.Context, req *health.HealthCheckRequest, reqId string) (*health.HealthCheckResponse, error)
	}
)

func NewHealthCheckService(healthCheckRepo *repository.HealthCheckRepoIn) HealthCheckServiceIn {
	return &healthCheckService{
		HealthCheckRepo: *healthCheckRepo,
	}
}

func (s *healthCheckService) Check(ctx context.Context, req *health.HealthCheckRequest, reqId string) (*health.HealthCheckResponse, error) {
	//logic
	var res health.HealthCheckResponse
	res.Status = "1"
	return &res, nil
}
