package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func main() {
	http.HandleFunc("/", getRoot)

	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":3333",
		TLSConfig: tlsConfig,
	}

	err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println(err.Error())
	}
}
