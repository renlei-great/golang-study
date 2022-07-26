package main

import (
	"gotour/ch22rpc/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	http.HandleFunc(rpc.DefaultRPCPath, func(writer http.ResponseWriter, request *http.Request) {
		conn, _, err := writer.(http.Hijacker).Hijack()
		if err != nil {
			log.Println("rpc hijacker:", err.Error())
			return
		}
		var connected = "200 Connected to json rpc"
		io.WriteString(conn, "HTTP/1.0 " + connected+ "\n\n")
		jsonrpc.ServeConn(conn)
	})

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalln("listen error:", e)

	}
	http.Serve(l, nil)
}

func test(){
	rpc.RegisterName("MathService", new(server.MathService))
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalln("listen error:", e)

	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("jsonrpc.Serce: accept:", err.Error())
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}
