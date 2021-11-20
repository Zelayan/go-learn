package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	// with tls
	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello. world!\n")
}
