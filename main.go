package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:2226")
	if err != nil {
		log.Fatal("could not open socket", "error", err)
	}

	log.Info("running on " + ln.Addr().String())

	for {
		handleOne(ln)
	}
}

func handleOne(ln net.Listener) {
	conn, err := ln.Accept()
	defer func() {
		if err := conn.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			log.Warn("could not close connection", "error", err)
		}
	}()
	if err != nil {
		log.Warn("could not accept connection", "error", err)
		return
	}
	bts, _, err := bufio.NewReader(conn).ReadLine()
	if err != nil {
		log.Warn("could not process connection", "error", err)
		return
	}

	args := strings.Fields(string(bts))
	if err := open(args...); err != nil {
		log.Warn("could not process request", "error", err)
		return
	}
	log.Info("opened", "args", args)
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
