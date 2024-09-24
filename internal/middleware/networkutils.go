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

	for _, iface := range interfaces {
		if !isValidInterface(iface) {
			continue
		}

		ip, err := getInterfaceIP(iface)
		if err != nil {
			continue
		}

		if ip != "" {
			return ip, nil
		}
	}
	return "", fmt.Errorf("could not find any non-loopback IP address")
}

func isValidInterface(iface net.Interface) bool {
	return iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0
}

func getInterfaceIP(iface net.Interface) (string, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ip := extractIP(addr)
		if ip != "" {
			return ip, nil
		}
	}
	return "", nil
}

func extractIP(addr net.Addr) string {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}

	if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
		return ip.String()
	}
	return ""
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
