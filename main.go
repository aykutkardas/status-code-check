package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getstatuscode(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.StatusCode, "|", url)
	}
	resp.Body.Close()
}

func operatecommand(args []string) (url string, path string, status bool) {
	argsLen := len(args)

	if argsLen == 1 && args[0] == "-h" {
		fmt.Println("-- HELP TEXT --") // TODO: Write help text and docs
		return
	} else if argsLen < 2 {
		fmt.Println("You must specify a file path or url. For more information, please use the '-h' flag.")
		return
	}

	flag := args[0]

	if flag == "-u" {
		return args[1], "", true
	} else if flag == "-f" {
		return "", args[1], true
	}

	return "", "", false
}

func main() {
	arguments := os.Args[1:]
	url, path, status := operatecommand(arguments)

	if status != true {
		return
	}

	if len(url) > 0 {
		getstatuscode(url)
	} else if len(path) > 0 {
		dat, err := ioutil.ReadFile(path)
		if err == nil {
			fixedURLList := strings.ReplaceAll(string(dat), "\r\n", "\n") // windows new line problem solved.
			urlList := strings.Split(fixedURLList, "\n")
			for _, url := range urlList {
				getstatuscode(url)
			}
		}
	}

}
