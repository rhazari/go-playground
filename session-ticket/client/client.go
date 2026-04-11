// client.go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// makeRequest sends an HTTP GET request and prints the response.
func makeRequest(client *http.Client) {
	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatalf("client: Get: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("client: ReadAll: %s", err)
	}
	log.Printf("Response: %s", body)
}

// SessionData is a storable format for the TLS session state
type SessionData struct {
	Ticket     []byte
	StateBytes []byte
}

// SaveSessionToFile serializes a ClientSessionState and writes it to the specified file
func SaveSessionToFile(filename string, cs *tls.ClientSessionState) error {
	fmt.Println("SaveSessionToFile:")
	ticket, state, err := cs.ResumptionState()
	if err != nil {
		return fmt.Errorf("failed to get resumption state: %w", err)
	}

	stateBytes, err := state.Bytes()
	if err != nil {
		return fmt.Errorf("failed to serialize session state: %w", err)
	}

	data := SessionData{
		Ticket:     ticket,
		StateBytes: stateBytes,
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("failed to encode session data with gob: %w", err)
	}

	return nil
}

// LoadSessionFromFile reads session data from a file and reconstructs the ClientSessionState
func LoadSessionFromFile(filename string) (*tls.ClientSessionState, error) {
	fmt.Println("LoadSessionFromFile:")
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var data SessionData
	enc := gob.NewDecoder(file)
	if err := enc.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode session data with gob: %w", err)
	}

	state, err := tls.ParseSessionState(data.StateBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse session state: %w", err)
	}
	// You do not unmarshal 'data.StateBytes' further using a tls-specific function.
    // The bytes in data.StateBytes are already in the format needed by NewResumptionState.

	// Reconstruct the final *tls.ClientSessionState using NewResumptionState
	cs, err := tls.NewResumptionState(data.Ticket, state)

	return cs, nil
}

// CustomSessionCache implements tls.ClientSessionCache
type CustomSessionCache struct {
    filename string
    // A mutex is required as the cache can be accessed concurrently by different goroutines.
    mu       sync.Mutex
}

// NewCustomSessionCache creates a new session cache tied to a file
func NewCustomSessionCache(filename string) *CustomSessionCache {
    return &CustomSessionCache{filename: filename}
}

// Get retrieves a ClientSessionState from the file
func (c *CustomSessionCache) Get(sessionKey string) (*tls.ClientSessionState, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

	fmt.Println("Get tls.ClientSessionState:")

    // Note: The sessionKey is typically the "host:port" of the server.
    // In this file-based example, we assume we only store one session per filename,
    // so we can ignore the sessionKey or use it to manage multiple files if needed.
    cs, err := LoadSessionFromFile(c.filename)
    if err != nil {
        // If loading fails (e.g., file not found), return false.
        return nil, false
    }
    return cs, true
}

// Put saves a ClientSessionState to the file
func (c *CustomSessionCache) Put(sessionKey string, cs *tls.ClientSessionState) {
    c.mu.Lock()
    defer c.mu.Unlock()

    if cs == nil {
        // If nil is passed, remove the cache entry (delete the file).
        os.Remove(c.filename)
        return
    }
	fmt.Println("Put tls.ClientSessionState:")

    // Save the session state to the file
    err := SaveSessionToFile(c.filename, cs)
    if err != nil {
        fmt.Printf("Error saving session to file: %v\n", err)
    }
}

func main() {
	// Load the server's self-signed certificate to trust it
	caCert, err := os.ReadFile("../server.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a session cache. Go's LRU cache is a great default choice.
	// sessionCache := tls.NewLRUClientSessionCache(32) // Cache up to 32 sessions

	// --- Use your custom file-based session cache instead of NewLRUClientSessionCache ---
	sessionCacheFile := "client_session_cache.gob"
	// Instantiate your custom file cache implementation
	fileCache := NewCustomSessionCache(sessionCacheFile)

	// Create the TLS configuration with the session cache
	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		ClientSessionCache: fileCache,
	}

	// Create an HTTP client that reuses the same TLS config
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// --- Step 1: Perform the first connection (full handshake) ---
	fmt.Println("--- First connection (full handshake) ---")
	makeRequest(client)

	// Wait for a moment to simulate a gap between connections
	fmt.Println("\nWaiting for 1 second before reconnecting...")
	time.Sleep(1 * time.Second)

	// --- Step 2: Perform the second connection (session resumption) ---
	fmt.Println("\n--- Second connection (session resumption) ---")
	makeRequest(client)
}