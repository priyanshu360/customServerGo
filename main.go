package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type serverStruct struct {
	*http.Server
}

func healthCheck(rw http.ResponseWriter, r *http.Request) {
	log.Println("Getting health check status")
	io.WriteString(rw, "Status:OK")
}

func (srv serverStruct) Start() error {
	log.Println("Server is running on Port :" + port)
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}

var port string

func main() {
	port = os.Args[1]
	myHandler := http.NewServeMux()

	myHandler.HandleFunc("/", healthCheck)
	var myServer serverStruct
	myServer.Server = &http.Server{
		Addr:           port,
		Handler:        myHandler,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(myServer.Start())
}
