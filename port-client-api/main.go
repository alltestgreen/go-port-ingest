package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alltestgreen/go-port-ingest/port-client-api/handler"
	"github.com/alltestgreen/go-port-ingest/port-client-api/service"
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
	portService := service.NewPortService()
	return handler.ServeEndpoints(portService)
}

// clean up resources before the application stopped
func shutdownApp() error {
	return nil
}
