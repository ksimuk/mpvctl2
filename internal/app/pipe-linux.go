//go:build linux

package app

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

const PIPE_LINUX = `/tmp/mpv-socket`

func ensureConnect(conn net.Conn) bool {
	fmt.Fprint(conn, `{"command": ["get_version"]}`)
	fmt.Fprint(conn, "\n")
	conn.SetReadDeadline(time.Now().Add(time.Second))
	_, err := bufio.NewReader(conn).ReadString('\n')
	return err == nil
}

func connectPipe() (net.Conn, error) {
	var err error
	for i := 1; i < 10; i++ {
		var conn net.Conn
		conn, err = net.Dial("unix", PIPE_LINUX)
		active := ensureConnect(conn)
		if err == nil && active {
			return conn, nil
		}
		closeSocket(conn)
		fmt.Println("retrying socket...")
		time.Sleep(time.Millisecond * 50)
	}

	return nil, err
}

func getPlaylistPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dirname, `.mpv_playlist.txt`)
}

func closeSocket(pipe net.Conn) {
	pipe.Close()
}
