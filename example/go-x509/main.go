package main

import (
	"crypto/rand"
	"crypto/x509"
)

func main() {
	x509.CreateCertificate(rand.Read)
}
