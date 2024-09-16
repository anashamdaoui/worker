package middleware

import (
	"fmt"
	"net"
	"os"
	"time"

	"math/rand"
)

/*
In Docker and Kubernetes environments, pass environment variable to the container that contain the IP address of the host machine or the container itself.
*/
func getIPFromEnv() (string, error) {
	ip := os.Getenv("HOST_IP")
	if ip == "" {
		return "", fmt.Errorf("HOST_IP not set")
	}
	return ip, nil
}

/*
Iterate over all the available network interfaces on the system and pick a non-loopback interface.
This method works for both VMs and Docker containers and does not rely on external connectivity
*/
func getLocalIPByInterface() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range interfaces {
		// Ignore down and loopback interfaces
		if i.Flags&net.FlagUp == 0 || i.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Skip loopback addresses
			if ip == nil || ip.IsLoopback() {
				continue
			}

			// Return the first non-loopback IP address found
			if ip.To4() != nil {
				return ip.String(), nil
			}
		}
	}
	return "", fmt.Errorf("could not find any non-loopback IP address")
}

// Find an available IP address for the worker by trying multiple methods
func GetLocalIP() (string, error) {
	// First try environment variable
	IP, err := getIPFromEnv()
	if err != nil {

		// Then try scanning network interfaces
		IP, err := getLocalIPByInterface()
		if err != nil {
			return "", err
		}

		return IP, nil
	}

	return IP, err
}

// Random port generator between min and max
func GetRandomPort(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min
}
