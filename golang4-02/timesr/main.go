package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	messages = make(chan string)
)

func main() {
	ln, err := net.Listen("tcp", "localhost:9002")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("Server is listening...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("send message:")
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		conn.Write([]byte(msg + "\n"))
		time.Sleep(1 * time.Second)
	}
}
