package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
)

// This function handles a request, and calls the correct function based on the path
func ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	path = req.URL.Path

}

func main() {
	// listen for requests
	fmt.Println("Hello!")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}
	srv := http.NewServeMux()
	srv.HandleFunc("/", ServeHTTP)
	fcgi.Serve(listener, srv)
}
