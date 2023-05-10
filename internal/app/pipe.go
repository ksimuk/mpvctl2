//go:build !linux && !windows

package app

import (
	"net"
)

func connectPipe() (net.Conn, error) {
	return nil, nil
}

func getPlaylistPath() string {
	return ""
}

func closeSocket(pipe net.Conn) {

}
