package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	mu          sync.RWMutex
	currentCert tls.Certificate
	caCertPool  *x509.CertPool
)

func loadCertificates(certPath, keyPath, caCertPath string) error {
	mu.Lock()
	defer mu.Unlock()

	// Load client certificate and key
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return err
	}
	currentCert = cert
	log.Println("Loading certs for https client...")

	// Load CA certificate
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		return err
	}
	newCaCertPool := x509.NewCertPool()
	newCaCertPool.AppendCertsFromPEM(caCert)
	caCertPool = newCaCertPool
	return nil
}

func getClientCertificate(info *tls.CertificateRequestInfo) (*tls.Certificate, error) {
	mu.RLock()
	defer mu.RUnlock()
	return &currentCert, nil
}

func main() {
	// Request /hello over port 8080 via the GET method
	// r, err := http.Get("http://localhost:8080/hello")

	// Request /hello over HTTPS port 8443 via the GET method
	// r, err := http.Get("https://localhost:8443/hello")

	// Create a CA certificate pool and add cert.pem to it
	// caCert, err := os.ReadFile("../cert.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// Read the key pair to create certificate
	// cert, err := tls.LoadX509KeyPair("../cert.pem", "../key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create a HTTPS client and supply the created CA pool and certificate
	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			RootCAs: caCertPool,
	// 			Certificates: []tls.Certificate{cert},
	// 		},
	// 	},
	// }

	// Initial load of certificates
	err := loadCertificates("../cert.pem", "../key.pem", "../cert.pem")
	if err != nil {
		log.Fatalf("Failed to load initial certificates: %v", err)
	}

	tlsConfig := &tls.Config{
		GetClientCertificate: getClientCertificate,
		RootCAs:              caCertPool, // RootCAs still needs to be set, potentially updated separately
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	// Request /hello via the created HTTPS client over port 8443 via GET
	r, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	log.Println("Before refreshing certs...")
	fmt.Printf("%s\n", body)

	// Simulate certificate refresh
	log.Println("Refreshing mTLS certificates dynamically...")
	err = loadCertificates("../cert.pem", "../key.pem", "../cert.pem") // Update the global variables
	if err != nil {
		log.Fatalf("Failed to load new certificates: %v", err)
	}

	// Request /hello via the created HTTPS client over port 8443 via GET
	r, err = client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err = io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}


// Generate certs
// openssl req -newkey rsa:2048 \
//   -new -nodes -x509 \
//   -days 3650 \
//   -addext "subjectAltName = DNS:localhost" \
//   -out cert.pem \
//   -keyout key.pem \
//   -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"
