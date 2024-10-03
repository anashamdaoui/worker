package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	grpcServer "worker/grpc/server"
	"worker/internal/config"
	"worker/internal/middleware"
	httpServer "worker/internal/server"
	"worker/internal/worker"

	"github.com/gorilla/mux"
)

func main() {

	// Load configuration
	config.LoadConfig("internal/config/config.json")

	// Initialize the logger with the configured log level
	middleware.InitLogger(config.AppConfig.LogLevel)
	logger := middleware.GetLogger()

	logger.Info("", "Starting worker ...")

	// Create a new router
	router := mux.NewRouter()

	// Channel to signal when the server is ready
	ready := make(chan struct{})

	// Create a new Worker instance
	err := worker.InitializeWorker()

	if err != nil {
		log.Fatal("Fatal Error: New worker instanciation failed.")
	}

	// Start the server in a separate goroutine
	srv := httpServer.StartHTTPServer(router, ready, worker.WorkerInstance.HTTPPort)

	// Wait for the server to signal readiness
	<-ready
	worker.Register()
	logger.Info("", "Server is ready to handle requests.")

	// Start GRPC server
	grpcServer.StartGRPCServer(worker.WorkerInstance.GRPCPort)

	// Set up signal handling for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	logger.Info("", "Shutting down worker...")

	if err := srv.Close(); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	logger.Info("", "Server exited properly")
}
