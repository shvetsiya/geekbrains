package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Enter your nickname:")
	reader := bufio.NewReader(os.Stdin)
	nick, err := reader.ReadString('\n')
	nick = strings.Trim(nick, " \r\n")

	go func() {
		io.Copy(os.Stdout, conn)
	}()
	io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit", conn.LocalAddr())
}

func read(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			conn.Close()
			os.Exit(0)
			return
		}

		fmt.Println(msg)
	}
}

func write(conn net.Conn, nick string) {
	for {
		reader := bufio.NewReader(os.Stdout)
		msg, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			break
		}
		nick = nick + "+" + strings.Trim(msg, " \r\n")
		conn.Write([]byte(nick))
	}
}
