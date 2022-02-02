package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type message struct {
	Сommand []string `json:"command"`
}

type playlistItem struct {
	Filename string `json:"filename"`
	Id       int    `json:"id"`
	Current  bool   `json:"current,omitempty"`
}

type response struct {
	Data  json.RawMessage `json:"data"`
	Error string          `json:"error"`
}

func getPlaylist(res response) []playlistItem {
	data := res.Data
	var playlist []playlistItem
	json.Unmarshal([]byte(data), &playlist)
	return playlist
}
func sendMessage(args []string) (*response, error) {
	conn, err := connectPipe()
	if err != nil {
		return nil, err
	}
	message := message{Сommand: args}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	fmt.Fprint(conn, string(messageJSON))
	fmt.Fprint(conn, "\n")

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return nil, err
	}
	var data response
	json.Unmarshal([]byte(msg), &data)

	return &data, nil
}

func savePlaylist(list []playlistItem) {
	var names []string
	current := false
	for _, item := range list {
		if item.Current {
			current = true
		}
		if current {
			names = append(names, item.Filename)
		}
	}
	res := strings.Join(names, "\n")
	err := ioutil.WriteFile(getPlaylistPath(), []byte(res), 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadPlaylist() []string {
	file, err := ioutil.ReadFile(getPlaylistPath())
	check(err)
	return strings.Split(string(file), "\n")
}

func run() {
	fmt.Println("WIP")
}

func Main() {
	args := os.Args[1:]
	switch args[0] {
	case "add":
		urls := args[1:]
		for _, url := range urls {
			fmt.Printf("Adding %s\n", url)
			_, err := sendMessage([]string{"loadfile", url, "append-play"})
			check(err)
		}
	case "play":
		_, err := sendMessage([]string{"set", "pause", "no"})
		check(err)
	case "pause":
		_, err := sendMessage([]string{"set", "pause", "yes"})
		check(err)
	case "next":
		_, err := sendMessage([]string{"playlist_next"})
		check(err)
	case "previous":
		_, err := sendMessage([]string{"playlist_prev"})
		check(err)
	case "playlist-count":
		res, err := sendMessage([]string{"get_property", "playlist/count"})
		check(err)
		fmt.Printf("Total %s\n", res.Data)
	case "playlist-pos":
		res, err := sendMessage([]string{"get_property", "playlist-pos"})
		check(err)
		fmt.Printf("Position %s\n", res.Data)
	case "playlist-remove":
		_, err := sendMessage([]string{"playlist-remove", os.Args[1]})
		check(err)
	case "playlist-clear":
		_, err := sendMessage([]string{"playlist-clear"})
		check(err)
	case "5s":
		_, err := sendMessage([]string{"seek", "5"})
		check(err)
	case "-5s":
		_, err := sendMessage([]string{"seek", "-5"})
		check(err)
	case "playlist":
		res, err := sendMessage([]string{"get_property", "playlist"})
		check(err)
		playlist := getPlaylist(*res)
		for _, item := range playlist {
			current := " "
			if item.Current {
				current = ">"
			}
			fmt.Printf("%s %d. %s\n", current, item.Id, item.Filename)
		}
	case "save-playlist":
		res, err := sendMessage([]string{"get_property", "playlist"})
		check(err)
		playlist := getPlaylist(*res)
		savePlaylist(playlist)

	case "load-playlist":
		urls := loadPlaylist()
		for _, url := range urls {
			fmt.Printf("Adding %s\n", url)
			_, err := sendMessage([]string{"loadfile", url, "append-play"})
			check(err)
		}
	case "start":
		run()
	default:
		fmt.Printf("%s not supported\n", args[0])
	}

}
