package network

import (
	"context"
	"grpc-boiler-plate-go/app/model/proto/health"
	"grpc-boiler-plate-go/app/presenter"
	"grpc-boiler-plate-go/env"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	grpcLog = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
)

func init() {
	grpclog.SetLoggerV2(grpcLog)
}

// InitGRPCServer : Init gRPC service as Server
func InitGRPCServer(dep *env.Dependency) *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(interceptor(dep)))
	presenter := presenter.GRPCPresenter(dep) // Inject dependencies
	health.RegisterHealthCheckServiceServer(srv, presenter)
	return srv
}

// interceptor : all incoming request into the service will be intercepted here, expecting any kind of handlers before
// the main application start processing the request
func interceptor(di *env.Dependency) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)

		if di.GRPCLogMode {
			log.Println("intercepted for audit", err)
			grpcLog.Infof("[Method]: - %s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
			grpcLog.Infof("[Request]: - %s\n", req)
			grpcLog.Infof("[Response]: - %s\n", resp)
		}

		return
	}
}
