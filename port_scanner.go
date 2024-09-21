// Port Scanner in Go
// Author: Ryan Feneley
// Date: September 2024

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// scan a single port
func scanPort(protocol, hostname string, port int) bool {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// Banner grabbing function to get service information
func grabBanner(hostname string, port int) string {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return ""
	}
	defer conn.Close()
	// Set read timeout
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	// Create buffer to read banner
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return ""
	}

	return string(buffer[:n])
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run portscanner.go <hostname> <startPort> <endPort>")
		os.Exit(1)
	}

	hostname := os.Args[1]
	startPort := 1
	endPort := 1024
	fmt.Sscanf(os.Args[2], "%d", &startPort)
	fmt.Sscanf(os.Args[3], "%d", &endPort)

	fmt.Printf("Scanning %s from port %d to %d...\n", hostname, startPort, endPort)

	for port := startPort; port <= endPort; port++ {
		open := scanPort("tcp", hostname, port)
		if open {
			fmt.Printf("Port %d is open\n", port)

			banner := grabBanner(hostname, port)
			if banner != "" {
				fmt.Printf("Banner for port %d: %s\n", port, banner)
			}
		}
	}
}
