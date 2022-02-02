//go:build windows

package app

import "gopkg.in/natefinch/npipe.v2"

const PIPE_WINDOWS = `\\.\pipe\mpv-pipe`

func connectPipe() (*PipeConn, error) {
	conn, err := npipe.Dial(PIPE_WINDOWS)
	if err != nil {
		return nil, err
	}
	return conn
}

func getPlaylistPath() string {
	return `C:\Tools\mpv.playlist`
}
