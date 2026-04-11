package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// generateKey generates a secure 32-byte key for use in session tickets.
func generateKey() [32]byte {
	var key [32]byte
	_, err := rand.Read(key[:])
	if err != nil {
		log.Fatalf("failed to generate session ticket key: %s", err)
	}
	return key
}

// rotateKeys manages the rotation of session ticket keys.
// The first key in the slice is used for encryption, while all keys are used for decryption.
// This allows clients with older tickets to still resume their sessions.
func rotateKeys() [][32]byte {
	// For this example, we will simply generate a new key on startup.
	// In a production environment, you would manage a slice of keys,
	// rotating them on a schedule (e.g., daily) and sharing them
	// across all server instances.
	newKey := generateKey()
	// You might also load previous keys from a secure, persistent storage.
	return [][32]byte{newKey}
}

func main() {
	go func() {
    	log.Println("Starting pprof server on http://localhost:6060/debug/pprof/")
    	log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// A handler for our server
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received connection from client")
		state := r.TLS
		if state != nil && state.DidResume {
			log.Println("Connection resumed a previous session!")
		} else {
			log.Println("New connection established.")
		}
		fmt.Fprintf(w, "Hello, TLS! DidResume: %v", state != nil && state.DidResume)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// Load the certificate and key
	cert, err := tls.LoadX509KeyPair("../server.crt", "../server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	// Create a TLS config and set the session ticket keys
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS13,
	}

	// Set the session ticket keys manually
	keys := rotateKeys()
	tlsConfig.SetSessionTicketKeys(keys)

	server := &http.Server{
		Addr:      "localhost:8443",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	log.Println("Starting server on https://localhost:8443")
	log.Fatal(server.ListenAndServeTLS("", ""))
}