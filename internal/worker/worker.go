package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"worker/internal/config"
	"worker/internal/middleware"

	"github.com/google/uuid"
)

type Worker struct {
	IP             string
	Port           int
	ID             string
	LastRegistered time.Time
}

var WorkerInstance Worker

func generateUUID() string {
	// Generate a new UUID
	id := uuid.New()
	return id.String()
}

func InitializeWorker() error {

	logger := middleware.GetLogger()
	logger.Debug("", "Creating NewWorker")

	ip, err := middleware.GetLocalIP()

	if err != nil {
		logger.Debug("", "Error: Failed to get local IP")
		// It is not critical for the worker to manage to get its local IP
	}

	port := -1
	if config.AppConfig.ServerPort != "" {
		port, _ = strconv.Atoi(config.AppConfig.ServerPort)
	} else {
		logger.Debug("", "GetRamdomPort")
		port = middleware.GetRandomPort(3060, 4000)
	}

	WorkerInstance.IP = ip
	WorkerInstance.Port = port
	WorkerInstance.ID = generateUUID()

	logger.Info("", "Woker initialized : \n\t - IP %s \n\t - Port %d \n\t - UUID %s", WorkerInstance.IP, WorkerInstance.Port, WorkerInstance.ID)

	logger.Debug("", "NewWorker instance created successfully")

	return nil
}

// Send a POST /register to the registry server with worker port in the body
func registerWorker(registryURL string, id string, port int, apiKey string) (*http.Response, error) {
	workerData := map[string]interface{}{
		"id":   id,
		"port": port,
	}

	jsonData, _ := json.Marshal(workerData)

	req, err := http.NewRequest("POST", registryURL+"/register", bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	// Add the required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %s", err)
	}

	defer resp.Body.Close()

	return resp, nil
}

func Register() {
	logger := middleware.GetLogger()
	logger.Debug("", "Registering to the registry-service %s", config.AppConfig.RegistryURI)

	isRegistered := false

	const retries = 4
	for i := 0; i < retries; i++ {
		resp, err := registerWorker(config.AppConfig.RegistryURI, WorkerInstance.ID, WorkerInstance.Port, config.AppConfig.APIKey)
		isRegistered = err == nil && resp.StatusCode == http.StatusOK
		if !isRegistered {
			// Log the error and implement a retry mechanism
			logger.Debug("", "Error registring worker : %v. Retrying...", err)
		} else {
			WorkerInstance.LastRegistered = time.Now()
			logger.Info("", "Worker registered successfully at %s", WorkerInstance.LastRegistered.String())
			break
		}
		time.Sleep(1000 * time.Millisecond) // Backoff
	}

	if !isRegistered {
		log.Fatalf("Fatal Error: worker did not manage to register... Exiting")
	}

}
