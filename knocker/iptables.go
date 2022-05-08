package knocker

import (
	"fmt"
	"os/exec"
	"syscall"
)

func checkIfPortAlreadyOpend(sourceAddress string, port string) (bool, error) {
	cmd := exec.Command("iptables", "-C", "knocker", "-s", sourceAddress, "-p", "tcp", "--dport", port, "-j", "ACCEPT")
	err := cmd.Start()
	if err != nil {
		return false, err
	}
	err = cmd.Wait()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus() == 0, nil
			}
		} else {
			return false, err
		}
	}
	fmt.Printf("lol\n")
	return false, nil
}

func openPortForAddress(sourceAddress string, port string) error {
	cmd := exec.Command("iptables", "-A", "knocker", "-s", sourceAddress, "-p", "tcp", "--dport", port, "-j", "ACCEPT")
	return cmd.Run()
}

func removeOpenPortForAddress(sourceAddress string, port string) error {
	cmd := exec.Command("iptables", "-D", "knocker", "-s", sourceAddress, "-p", "tcp", "--dport", port, "-j", "ACCEPT")
	return cmd.Run()
}
