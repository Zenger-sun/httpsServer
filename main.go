package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

const (
	WEB_DIR = "./web"
	ADDR = "127.0.0.1:8001"
	CERT = "./cert/server.pem"
	PRIV = "./cert/server.key"
)

func main() {
	log.Println("start https server on", ADDR)

	//httpServer()
	echoServer()
}

func httpServer()  {
	http.Handle("/", http.FileServer(http.Dir(WEB_DIR)))
	err := http.ListenAndServeTLS(ADDR, CERT, PRIV, nil)
	if err != nil {
		panic(err)
	}
}

func echoServer() {
	const pack_len = 4
	var err error

	certs := make([]tls.Certificate, 1)
	certs[0], err = tls.LoadX509KeyPair(CERT, PRIV)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		Certificates: certs,
	}

	listener, _ := tls.Listen("tcp", ADDR, config)
	defer listener.Close()

	for {
		conn, _ := listener.Accept()

		go func(conn net.Conn) {
			defer conn.Close()
			b := make([]byte, 32 * 1024)

			for {
				_, err := io.ReadAtLeast(conn, b, pack_len)
				switch err {
				case nil:
				default:
					fmt.Println(err)
					return
				}

				l := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
				pack := b[pack_len:pack_len+l]
				fmt.Println(string(pack))

				conn.Write(pack)
			}
		}(conn)
	}
}
