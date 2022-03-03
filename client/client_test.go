package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const (
	URL  = "https://127.0.0.1:8001/index.html"
	ADDR = "3.15.3.128:8001"
	CA   = "../cert/ca.crt"
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

	client := &http.Client{Transport: tr}

	r, err := client.Get(URL)
	if err != nil {
		fmt.Errorf("get %s error", URL)
		panic(err)
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("get %s error", URL)
		panic(err)
	}

	fmt.Println(string(buf))
}

func Test_echoServer(t *testing.T) {
	rootCa, err := os.ReadFile(CA)
	if err != nil {
		panic("failed to read root certificate")
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootCa)
	if !ok {
		panic("failed to parse root certificate")
	}

	tlsConf := &tls.Config{
		RootCAs:    roots,
		NextProtos: []string{"quic-echo-example"},
	}

	conn, err := tls.Dial("tcp", ADDR, tlsConf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("message%d", i)
		buff := new(bytes.Buffer)
		binary.Write(buff, binary.LittleEndian, int32(len(msg)))
		buff.Write([]byte(msg))

		_, err = conn.Write(buff.Bytes())
		if err != nil {
			panic(err)
		}

		pack := make([]byte, buff.Len() - 4)
		io.ReadFull(conn, pack)
		fmt.Println(string(pack))
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

	client := &http.Client{Transport: tr}

	for i := 0; i < b.N; i++ {
		_, err := client.Get(URL)
		if err != nil {
			fmt.Errorf("get %s error", URL)
			panic(err)
		}
	}
}

func Benchmark_echoServer(b *testing.B) {
	for i := 0; i <  b.N; i++ {
		rootCa, err := os.ReadFile(CA)
		if err != nil {
			panic("failed to read root certificate")
		}

		roots := x509.NewCertPool()
		ok := roots.AppendCertsFromPEM(rootCa)
		if !ok {
			panic("failed to parse root certificate")
		}

		tlsConf := &tls.Config{
			RootCAs:    roots,
			NextProtos: []string{"quic-echo-example"},
		}

		conn, err := tls.Dial("tcp", ADDR, tlsConf)
		if err != nil {
			panic(err)
		}

		msg := fmt.Sprintf("message%d", i)
		buff := new(bytes.Buffer)
		binary.Write(buff, binary.LittleEndian, int32(len(msg)))
		buff.Write([]byte(msg))

		_, err = conn.Write(buff.Bytes())
		if err != nil {
			panic(err)
		}

		pack := make([]byte, buff.Len() - 4)
		io.ReadFull(conn, pack)
	}
}
