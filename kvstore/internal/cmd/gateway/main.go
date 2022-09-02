package main

import (
	"context"
	"fmt"
	"github.com/becosuke/golang-examples/kvstore/internal/registry/injection"
	"github.com/becosuke/golang-examples/kvstore/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
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

	idleConnsClosed := make(chan struct{})
	go func() {
		defer close(idleConnsClosed)
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
		<-sigCh
		logger.Info("received SIGTERM, system is going gracefully shutdown")
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Error("received error on gracefully shutdown", zap.Error(err))
		}
		logger.Info("http server has completed gracefully shutdown")
	}()
	logger.Info("http server trys to listen", zap.Int("port", config.HttpPort))
	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("http server failed to listen", zap.Error(err))
		return exitError
	}
	<-idleConnsClosed

	cancel() // does it need this?
	return exitOK
}
