package server

import (
	"log"
	"net/http"
	"strconv"
	"worker/internal/middleware"

	"github.com/gorilla/mux"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetRequestIDFromContext(r.Context())
	logger := middleware.GetLogger()
	logger.Debug(requestID, "Handling /healthcheck request")

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Worker is healthy")); err != nil {
		logger.Debug(requestID, "Error writing response: %v", err)
	}
}

// StartServer starts the HTTP server for the registry service and returns the server instance.
func StartHTTPServer(router *mux.Router, ready chan struct{}, port int) *http.Server {

	router.Use(middleware.RequestID)        // Add Request ID middleware
	router.Use(middleware.LoggerMiddleware) // Add Logger middleware
	router.Use(middleware.AuthMiddleware)   // Add Auth middleware

	router.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")

	portStr := strconv.Itoa(port)
	srv := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
	}

	log.Printf("Starting HTTP server on port : %d...", port)
	close(ready) // Signal that the server is ready
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server stopped: %v", err)
		}
	}()

	return srv
}
