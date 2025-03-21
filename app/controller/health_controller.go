package controller

import (
	"context"
	"fmt"
	"grpc-boiler-plate-go/app/model/proto/health"
	"grpc-boiler-plate-go/app/service"
	"grpc-boiler-plate-go/helpers"
	"grpc-boiler-plate-go/pkg/runtimekit"

	guuid "github.com/google/uuid"
)

type HealthCheckController struct {
	HealthCheckService service.HealthCheckServiceIn
	health.UnimplementedHealthCheckServiceServer
}

// NewUserController initializes the UserController
func NewHealthCheckController(healthCheckService *service.HealthCheckServiceIn) *HealthCheckController {
	return &HealthCheckController{
		HealthCheckService: *healthCheckService,
	}
}

func (c *HealthCheckController) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	requestId := guuid.New()
	logStart := helpers.LogRequest(ctx, req.String(), requestId.String(), runtimekit.FunctionName())
	fmt.Println(logStart)

	response, err := c.HealthCheckService.Check(ctx, req, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String(), runtimekit.FunctionName())
		return &health.HealthCheckResponse{
			Status: "0",
		}, nil
	}

	logStop := helpers.LogResponse(ctx, response.String(), requestId.String(), runtimekit.FunctionName())
	fmt.Println(logStop)

	return response, nil
}
