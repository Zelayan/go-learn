package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	// with tls
	//log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
	// Create a CA certificate pool and add cert.pem to i
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	tlsConfig.BuildNameToCertificate()

	// Create a Server instance to listen on port 8443 with the TLS config
	server := http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}
	// Listen to HTTPS connections with the server certificate and wait
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello. world!\n")
}
