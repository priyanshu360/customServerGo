package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func myHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Server is running")
}

func healthCheck(rw http.ResponseWriter, r *http.Request) {
	log.Println("Server is running on Port", port)
	io.WriteString(rw, "Status:OK")
}

var port string

func main() {
	port = os.Args[1]
	myHandler := http.NewServeMux()

	myHandler.HandleFunc("/", healthCheck)

	myServer := &http.Server{
		Addr:           port,
		Handler:        myHandler,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(myServer.ListenAndServe())
}
