package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultTimeout = 10 * time.Second

func main() {
	timeout := flag.Duration("timeout", defaultTimeout, "timeout to connect")
	flag.Parse()

	host, port := flag.Arg(0), flag.Arg(1)
	if host == "" || port == "" {
		log.Fatal("server address is not valid")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	addr := net.JoinHostPort(host, port)

	tc := NewTelnetClient(addr, *timeout, os.Stdin, os.Stdout)

	if err := tc.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	defer tc.Close()

	go func() {
		defer cancel()

		if err := tc.Send(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}()

	go func() {
		defer cancel()

		if err := tc.Receive(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}()

	<-ctx.Done()
}
