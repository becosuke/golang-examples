package main

import (
	"context"
	"fmt"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/injection"
	"github.com/becosuke/golang-examples/kvstore/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
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
	os.Exit(run())
}

func run() int {
	in := injection.NewInjection(serviceName, version)
	config := in.InjectConfig()
	logger := in.InjectLogger()
	defer logger.Sync()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterKVStoreServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", config.GrpcPort), opts)
	if err != nil {
		logger.Error("failed to register handler", zap.Error(err))
		return exitError
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", config.HttpPort),
		Handler: mux,
	}
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return httpServer.ListenAndServe() })
	logger.Info("httpServer listen", zap.Int("port", config.HttpPort))

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
		logger.Info("http server is going gracefully shutdown")
		httpServer.Shutdown(ctx)
		logger.Info("http server has completed gracefully shutdown")
	}()

	cancel()
	if err := wg.Wait(); err != nil {
		logger.Error("received error", zap.Error(err))
		return exitError
	}

	return exitOK
}
