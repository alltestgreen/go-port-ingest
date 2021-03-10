package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/alltestgreen/go-port-ingest/port-domain-service/server"
	"github.com/alltestgreen/go-port-ingest/proto"
)

func main() {
	errCh := startApp()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		log.Printf("server encountered error. initiating app shutdown: %v", err)
	case <-quit:
		log.Printf("os termination signal received. initiating app shutdown...")
	}

	err := shutdownApp()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("application shutdown complete")
}

func startApp() chan error {
	errCh := make(chan error)
	go func() {
		if err := initService(); err != nil {
			errCh <- err
		}
	}()
	return errCh
}

func initService() error {
	address := ":4040"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterPortServiceServer(srv, server.NewServer())

	log.Printf("gRPC server started listening on (%s)", address)

	return srv.Serve(listener)
}

// clean up resources before the application stopped
func shutdownApp() error {
	return nil
}
