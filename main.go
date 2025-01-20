// go_crypto_rule_tlsversion.go

// This Go file demonstrates best practices for enforcing TLS 1.3
// in a Go server, ensuring compliance with modern security standards.

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"log"
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

func generateRSAKey() {
	// Generate an RSA key with a minimum recommended size of 2048 bits
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated RSA Key with %d bits\n", key.Size()*8)
}

// Notes for RSA Key Strength:
// - Avoid generating RSA keys with less than 2048 bits to comply with modern security standards.
// - Keys of insufficient strength (e.g., 1024 bits) are deprecated by NIST and may soon be vulnerable due to advances in computing power.
// - Always validate the generated key size and store it securely.
