package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

const PIPE_LINUX = `/tmp/mpv-socket`

func connectPipe() (net.Conn, error) {
	var err error
	fmt.Println("1")
	for i := 1; i < 10; i++ {
		var conn net.Conn
		conn, err = net.Dial("unix", PIPE_LINUX)
		if err == nil {
			return conn, nil
		}
		fmt.Println("retrying socket...")
		time.Sleep(time.Second)
	}

	return nil, err
}

func getPlaylistPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dirname, `mpv_playlist.txt`)
}
