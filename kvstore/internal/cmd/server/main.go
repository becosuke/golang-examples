package main

import (
	"context"
	"github.com/becosuke/golang-examples/kvstore/internal/pb"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/injection"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	in := injection.NewInjection()
	grpcServer := in.InjectGrpcServer()
	kvstoreServiceServer := in.InjectKVStoreServiceServer()
	pb.RegisterKVStoreServiceServer(grpcServer, kvstoreServiceServer)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return grpcServer.Serve(grpcLn) })

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		logger.Info("received SIGTERM, exiting server gracefully")
	case <-ctx.Done():
	}

	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)
		logger.Info("gracefully shutdown gRPC server")
		grpcServer.GracefulStop()
		logger.Info("completed to gracefully shutdown gRPC server")
	}()
}
