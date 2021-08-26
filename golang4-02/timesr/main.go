package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGUSR1)
	cfg := net.ListenConfig{
		KeepAlive: time.Minute,
	}

	listener, err := cfg.Listen(ctx, "tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	log.Println("im started!")

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		} else {
			wg.Add(1)
			go handleConn(ctx, conn, wg)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("done")
			listener.Close()
			wg.Wait()
			log.Println("exit")
			return
		}
	}

}

func handleConn(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
	// every second the server sends its current time to a client
	tck := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-tck.C:
			fmt.Fprintf(conn, "now: %s\n", t)
		}
	}
}
