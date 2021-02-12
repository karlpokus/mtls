# mtls
A simple implementation of mutual tls in go, including generating certs.

# usage

gen certs

````bash
# server
$ go run cert.go -ca --host localhost --dest cert/server
# client
$ go run cert.go -ca --host localhost --dest cert/client -client
# optinally check cert contents
$ openssl x509 -in <file> -text
````

run thangs

````bash
$ go run server.go
$ go run client.go
````

# test

Use a mix of the `--key`, `--cert` and `--cacert` flags to test w curl.

proper error messages from server
- client missing host certs: remote error: tls: unknown certificate authority
- client missing client certs: tls: client didn't provide a certificate
- client using server certs: tls: failed to verify client certificate: x509: certificate signed by unknown authority
