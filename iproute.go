package iproute

import (
	"errors"
	"os/exec"
	"strings"
)

// LocalSrcIP gets the local source IP used for packets going to destination IP.
func LocalSrcIP(dstIP string) (string, error) {
	cmd := exec.Command("ip", "-o", "route", "get", dstIP)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(strings.TrimRight(string(stdoutStderr),"\n"))
	}

	parts := strings.Fields(string(stdoutStderr))
	for index, val := range parts {
		if val == "src" {
			devName := parts[index+1]
			return devName, nil
		}
	}

	return "", nil
}
