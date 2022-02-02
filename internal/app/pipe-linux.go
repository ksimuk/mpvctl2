package app

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

const PIPE_LINUX = `/tmp/mpv-socket`

func connectPipe() (net.Conn, error) {
	conn, err := net.Dial("unix", PIPE_LINUX)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getPlaylistPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dirname, `mpv_playlist.txt`)
}
