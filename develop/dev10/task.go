package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var host, port string
var timeout time.Duration

func main() {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println(net.JoinHostPort(host, port))
		panic(err)
	}
	defer conn.Close()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Read data
	go func() {
		reader := bufio.NewReader(conn)
		for {
			data, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err.Error())
			}
			fmt.Println(data)
		}
	}()

	// Write data
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := scanner.Text()
			_, err := fmt.Fprintln(conn, data)
			if err != nil {
				return
			}
		}
	}()
	<-sigCh
}

func init() {
	timeoutPtr := flag.Duration("timeout", 10*time.Second, "timeout duration")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		panic(errors.New("empty host or port"))
	}
	host = args[0]
	port = args[1]
	timeout = *timeoutPtr
}
