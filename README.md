# knocker
A simple port knocking server that uses iptables.

# Installation

Build the golang app with `go build -o knock`. You can run the binary with `sudo ./knock` (the server needs root privileges because he creates iptables rules)

## Iptables
You have to create the following iptables rules / chains:
```
-N knocker # this chain will be used for port releases by knocker
-A INPUT -j knocker
-A INPUT -p tcp --dport 9999 -j DROP # drop all packets to your ports that are not released by knocker.
```
Please make sure that the knock_listen_address port is opened in the firewall.

# Configuration
Replace the values with your needs.
```json
[
  {
    "knock_listen_address": "127.0.0.1:1234 (where the knock will be expected)",
    "open_port": "9999 (the port that is released for the knocking user)",
    "ttl": 60
  }
]
```
After ttl seconds the port will be blocked again.

# How to use?

Open a tcp connection to the `knock_listen_address`. The tcp connection will be closed after the establishment from the server. The open_port is now accessable by your ip address for ttl seconds.