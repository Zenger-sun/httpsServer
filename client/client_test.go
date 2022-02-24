package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"testing"
)

const (
	URL = "https://127.0.0.1:8001/"
	CA = "../cert/ca.crt"
)

func Test_httpsServer(t *testing.T) {
	rootCa, err := os.ReadFile(CA)
	if err != nil {
		panic("failed to read root certificate")
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootCa)
	if !ok {
		panic("failed to parse root certificate")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: roots,
		},
	}

	client := &http.Client{Transport:tr}

	_, err = client.Get(URL)
	if err != nil {
		fmt.Errorf("get %s error", URL)
		panic(err)
	}
}

func Benchmark_httpsServer(b *testing.B) {
	rootCa, err := os.ReadFile(CA)
	if err != nil {
		panic("failed to read root certificate")
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootCa)
	if !ok {
		panic("failed to parse root certificate")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: roots,
		},
	}

	client := &http.Client{Transport:tr}

	for i := 0; i < b.N; i++ {
		_, err := client.Get(URL)
		if err != nil {
			fmt.Errorf("get %s error", URL)
			panic(err)
		}
	}
}
