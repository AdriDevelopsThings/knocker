package knocker

import (
	"fmt"
	"net"
)

func handle(listen net.Listener, openPort string, ttl int64) {
	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Printf("Error while accepting connection on %q: %v\n", listen.Addr().String(), err)
			continue
		}
		fmt.Printf("Client %q knocked.\n", connection.RemoteAddr().String())
		go connectionCallback(connection, openPort, ttl)
	}
}

func StartServer(listenAddress string, openPort string, ttl int64) (net.Listener, error) {
	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Waiting for client connection on %q.\n", listenAddress)
	go handle(listen, openPort, ttl)
	return listen, nil
}

func StartServers() ([]net.Listener, error) {
	servers := make([]net.Listener, 0)
	for _, element := range ports {
		server, err := StartServer(element.ListenAddress, element.OpenPort, int64(element.TTL))
		if err != nil {
			return servers, err
		}
		servers = append(servers, server)
	}
	return servers, nil
}
