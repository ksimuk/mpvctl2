package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var cmd *exec.Cmd

const (
	INSTANCE_PORT = 9293
)

func listen(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go func() {
			conn.Write([]byte("I am alive"))
			conn.Close()
		}()
	}
}

func run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", INSTANCE_PORT))
	fmt.Println("listen")
	if err != nil {
		fmt.Println("already running")
		os.Exit(1) // already running
	}
	go listen(listener)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		kill()
		os.Exit(1)
	}()

	for {
		cmd := exec.Command(`C:\Users\maxim\scoop\apps\mpv\current\mpv.exe`, `--profile=service-windows`)
		cmd.Run()
		cmd.Wait()
		time.Sleep(5 * time.Second)
	}
}

func kill() error {
	if cmd == nil {
		return nil
	}
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(cmd.Process.Pid))
	kill.Stderr = os.Stderr
	kill.Stdout = os.Stdout
	return kill.Run()
}
