package knocker

import (
	"fmt"
	"net"
	"strings"
)

func connectionCallback(connection net.Conn, openPort string, ttl int64) {
	remoteAddr := strings.Split(connection.RemoteAddr().String(), ":")[0] + "/32"
	alreadyExists, err := checkIfPortAlreadyOpend(remoteAddr, openPort)
	if err != nil {
		fmt.Printf("Error while checking if port already opended for %q: %v\n", remoteAddr, err)
		return
	}
	if !alreadyExists {
		err = openPortForAddress(remoteAddr, openPort)
		if err != nil {
			fmt.Printf("Error while creating iptables rule for %q: %v\n", remoteAddr, err)
			return
		}
		addTTL(remoteAddr, openPort, ttl)
		fmt.Printf("Port %q opended for %q.\n", openPort, remoteAddr)
	}
	connection.Close()
}
