package main

import (
	"context"
	"fmt"
	"github.com/becosuke/golang-examples/kvstore/internal/pb"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/injection"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	exitOK int = iota
	exitError
)

var (
	serviceName string
	version     string
)

func main() {
	os.Exit(process())
}

func process() int {
	in := injection.NewInjection(serviceName, version)
	config := in.InjectConfig()
	logger := in.InjectLogger()
	defer logger.Sync()

	grpcServer := in.InjectGrpcServer()
	controller := in.InjectController()
	pb.RegisterKVStoreServiceServer(grpcServer, controller)
	reflection.Register(grpcServer)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcLn, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		logger.Error("failed to listen grpc port", zap.Error(err))
		return exitError
	}

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return grpcServer.Serve(grpcLn) })
	logger.Info("grpcServer listen", zap.Int("port", config.GrpcPort))

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		logger.Info("received SIGTERM, system is going gracefully shutdown")
	case <-ctx.Done():
	}

	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)
		logger.Info("grpc server is going gracefully shutdown")
		grpcServer.GracefulStop()
		logger.Info("grpc server has completed gracefully shutdown")
	}()

	cancel()
	if err := wg.Wait(); err != nil {
		logger.Error("received error", zap.Error(err))
		return exitError
	}

	return exitOK
}
