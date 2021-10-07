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

	"github.com/getlantern/systray"
)

var cmd *exec.Cmd
var shouldRun bool = true

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
	go systray.Run(onReady, onExit)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		kill()
	}()

	for {
		if !shouldRun {
			break
		}
		pathMpv, _ := exec.LookPath("mpv.exe")
		fmt.Printf("Path %s", pathMpv)
		cmd = exec.Command(pathMpv, `--profile=service-windows`)
		cmd.Run()
		cmd.Wait()
		time.Sleep(5 * time.Second)
	}
}

func kill() {
	shouldRun = false
	fmt.Println(cmd)
	if cmd != nil {
		kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(cmd.Process.Pid))
		kill.Run()
		kill.Wait()
	}
	os.Exit(0)
}
