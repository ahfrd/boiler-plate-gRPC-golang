package presenter

import (
	"grpc-boiler-plate-go/app/controller"
	"grpc-boiler-plate-go/app/repository"
	"grpc-boiler-plate-go/app/service"
	"grpc-boiler-plate-go/env"
)

func GRPCPresenter(di *env.Dependency) *controller.HealthCheckController {
	healthCheckRepo := repository.NewHealthCheckRepository(di)
	healthCheckService := service.NewHealthCheckService(&healthCheckRepo)
	healthCheckController := controller.NewHealthCheckController(&healthCheckService)
	return healthCheckController
}
