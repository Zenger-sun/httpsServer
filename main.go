package main

import (
	"log"
	"net/http"
	"os"
)

const (
	WEB_DIR = "./web"
	ADDR = "127.0.0.1:8001"
	CERT = "./cert/server.pem"
	PRIV = "./cert/server.key"
)

var staticHandler = http.FileServer(http.Dir(WEB_DIR))

func main() {
	log.Println("start https server on", ADDR)

	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServeTLS(ADDR, CERT, PRIV, nil)
	if err != nil {
		panic(err)
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		staticHandler.ServeHTTP(w, r)
		return
	}

	f, err := os.ReadFile(WEB_DIR+"/index.html")
	if err == nil {
		w.Write(f)
		return
	}

	w.Write([]byte("<h1>404 not found!</h1>"))
}
