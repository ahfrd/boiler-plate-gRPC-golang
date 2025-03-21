package repository

import "grpc-boiler-plate-go/env"

type (
	healthCheckRepo struct {
		di *env.Dependency
	}

	HealthCheckRepoIn interface {
		GetHealth() string
	}
)

func NewHealthCheckRepository(di *env.Dependency) HealthCheckRepoIn {
	return &healthCheckRepo{
		di: di,
	}
}

func (r *healthCheckRepo) GetHealth() string {
	return r.di.Params.Database.MySQLDB.DBURL
}
