package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type serverStruct struct {
	*http.Server
}

// var port = os.Getenv("PORT")
// var host = os.Getenv("HOST")
var address = "0.0.0.0" + ":" + "1323"

func HealthCheck(rw http.ResponseWriter, r *http.Request) {
	log.Println("Getting health check status")
	io.WriteString(rw, "Status:OK")
}

func (srv serverStruct) Start() error {
	log.Println("Server is running on " + address)
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}

var myHandler = http.NewServeMux()

func NewServerStruct() serverStruct {
	myHandler.HandleFunc("/", HealthCheck)
	myServer := serverStruct{
		Server: &http.Server{
			Addr:           address,
			Handler:        myHandler,
			ReadTimeout:    20 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
	return myServer
}

func main() {
	myServer := NewServerStruct()
	log.Fatal(myServer.Start())
}
