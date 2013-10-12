package main

import (
"fmt"
 "net"
 "net/http"
 "net/http/fcgi"
)

func ServeHTTP(resp http.ResponseWriter, req *http.Request) {
 resp.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))
}

func main() {
  fmt.Println("Hello!")
 listener, err := net.Listen("tcp", ":8080")
 if (err != nil) {
    fmt.Println("Error!")
    fmt.Println(err)
 }
 srv := http.NewServeMux()
 srv.HandleFunc("/", ServeHTTP)
 fcgi.Serve(listener, srv)
}