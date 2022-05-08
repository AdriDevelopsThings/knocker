package knocker

import (
	"fmt"
	"os"
	"time"
)

type TTLInformation struct {
	remoteAddr string
	port       string
	ttl        int64
}

var ttls []TTLInformation = make([]TTLInformation, 0)

func addTTL(remoteAddr string, port string, ttl int64) {
	ttls = append(ttls, TTLInformation{remoteAddr: remoteAddr, port: port, ttl: time.Now().Unix() + ttl})
}

func checkTTLs() {
	now := time.Now().Unix()
	for index, element := range ttls {
		if element.ttl < now {
			err := removeOpenPortForAddress(element.remoteAddr, element.port)
			if err != nil {
				fmt.Printf("Error while removing open port %q for %q: %v\n", element.port, element.remoteAddr, err)
				os.Exit(1)
				return
			}
			ttls[index] = ttls[len(ttls)-1]
			ttls = ttls[:len(ttls)-1]
			fmt.Printf("Removed open port %q for %q.\n", element.port, element.remoteAddr)
		}
	}
}

func StopAllTTLs() {
	for _, element := range ttls {
		removeOpenPortForAddress(element.remoteAddr, element.port)
		fmt.Printf("Removed open port %q for %q.\n", element.port, element.remoteAddr)
	}
	ttls = make([]TTLInformation, 0)
}

func CheckTTLsEverySeconds() {
	for {
		if len(ttls) > 0 {
			checkTTLs()
		}
		time.Sleep(2)
	}
}
