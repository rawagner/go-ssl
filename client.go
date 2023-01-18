package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	rootCAs := x509.NewCertPool()
	data, err := os.ReadFile("ca.crt")
	if err == nil {
		rootCAs.AppendCertsFromPEM(data)
	} else {
		fmt.Println(err.Error())
	}

	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println(err.Error())
	}

	tlsConfig := &tls.Config{
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{cert},
	}

	httpClient := &http.Client{Transport: &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: tlsConfig,
	}}

	resp, err := httpClient.Get("https://127.0.0.1:3333")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.StatusCode)
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(result))
}
