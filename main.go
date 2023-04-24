package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	ln, err := net.Listen("tcp", ":2226")
	if err != nil {
		log.Fatal("could not open socket", "error", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Warn("could not accept connection", "error", err)
			_ = conn.Close()
			continue
		}
		log.Info("got conn")
		bts, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			log.Warn("could not process connection", "error", err)
			_ = conn.Close()
			continue
		}

		args := strings.Fields(string(bts))
		log.Info("got args", "args", args)
		if err := open(args...); err != nil {
			log.Warn("could not process request", "error", err)
			_ = conn.Close()
			continue
		}
		_ = conn.Close()
	}
}

func open(args ...string) error {
	switch runtime.GOOS {
	case "darwin":
	case "linux":
		if d := os.Getenv("DISPLAY"); d == "" {
			return fmt.Errorf("DISPLAY not set")
		}
	default:
		return fmt.Errorf("unsupported OS")
	}
	return exec.Command("/usr/bin/open", args...).Run()
}
