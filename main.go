package main

import (
	"flag"
	"log"
	"net"
	"net/http"
)

var port *string

func init() {
	port = flag.String("port", "4000", "port to serve on")
}

func main() {
	flag.Parse()
	handler := http.FileServer(http.Dir("./static"))

	addr := net.JoinHostPort("127.0.0.1", *port)
	log.Printf("Serving on %s\n", addr)

	http.ListenAndServe(addr, handler)
}
