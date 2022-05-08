package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/adridevelopsthings/knocker/knocker"
)

func main() {
	err := knocker.ReadConfig()
	if err != nil {
		fmt.Printf("Error while reading config: %v\n", err)
		os.Exit(1)
	}
	servers, err := knocker.StartServers()
	if err != nil {
		fmt.Printf("Error while starting servers: %v\n", err)
		os.Exit(1)
	}
	setupCloseHandler(servers)
	knocker.CheckTTLsEverySeconds()
}

func setupCloseHandler(servers []net.Listener) {
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Printf("Shutting down.")
		knocker.StopAllTTLs()
		for _, server := range servers {
			server.Close()
		}
		os.Exit(0)
	}()
}
