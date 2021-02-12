package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	http.HandleFunc("/hello", hi)
	cert, err := tls.LoadX509KeyPair("cert/server/cert.pem", "cert/server/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	ca, err := ioutil.ReadFile("cert/client/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)
	server := &http.Server{
		Addr:      ":8989",
		TLSConfig: &tls.Config{
  		ClientCAs:  pool,
  		ClientAuth: tls.RequireAndVerifyClientCert,
      Certificates: []tls.Certificate{cert},
  	},
	}
  log.Fatal(server.ListenAndServeTLS("", ""))
}
