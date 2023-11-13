package main

import (
	"log"
	"net"
)

const PORT = "3000"

func handleConection(conn net.Conn) {
	conn.Write([]byte("Hello world"))
}

func main() {
	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Printf("Server has been started at port %s! 🚀\n", PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConection(conn)
	}
}
