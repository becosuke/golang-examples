package main

import (
	"github.com/becosuke/golang-examples/kvstore/internal/protogen"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/injector"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	nodeServer := newNodeServer

	nodeService, err := node.New()
	if err != nil {
		logger.Error("failed to setup node client", zap.Error(err))
		return exitError
	}

	// Create new gRPC server. gRPC server handles core function this service.
	grpcServer, err := grpc.New(nodeService, &grpc.Config{
		DDServiceName:        serviceName,
		DDStatsdClient:       ddStatsdClient,
		SentryClient:         sentryClient,
		AuthorityKeyRegistry: authorityClient,
		Logger:               logger,
	})
	if err != nil {
		logger.Error("failed to setup new gRPC server", zap.Error(err))
		return exitError
	}

	grpc.New
	protogen.RegisterNodeServiceServer()
	grpc.ServiceRegistrar()
	protogen.NodeServiceServer()
	in := injector.NewInjector()
	in.InjectServer().Serve()
}
