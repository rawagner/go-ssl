# go-ssl
`go run create-certs.go` to generate CA, client and server certificates for MTLS

`go run main.go` to run server which will use server cert and configures MTLS

`go run client.go` to create a GET request for server with client cert