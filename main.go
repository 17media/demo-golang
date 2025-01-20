package test

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	// Create a custom TLS configuration enforcing TLS 1.3
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	// Configure the HTTP server to use the TLS configuration
	httpServer := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	// Define a simple handler for testing purposes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS 1.3!")
	})

	fmt.Println("Starting server on https://localhost:8443")

	// Start the server with TLS
	err := httpServer.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

// Notes:
// - Ensure you have valid "server.crt" and "server.key" files.
// - Always prefer using TLS 1.3 to enforce modern cryptographic standards.
// - Deprecated versions like TLS 1.0 and TLS 1.1 should be avoided to prevent vulnerabilities.
