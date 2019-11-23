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

func operatecommand(args []string) (url, path string, status bool) {
	argsLen := len(args)

	type commandList struct {
		HELP string
		PATH string
		URL  string
	}

	command := commandList{HELP: "-h", PATH: "-f", URL: "-u"}

	if argsLen == 1 && args[0] == command.HELP {
		fmt.Println("-- HELP TEXT --") // TODO: Write help text and docs
		return
	} else if argsLen < 2 {
		fmt.Println("You must specify a file path or url. For more information, please use the '-h' flag.")
		return
	}

	flag := args[0]

	if flag == command.URL {
		return args[1], "", true
	} else if flag == command.PATH {
		return "", args[1], true
	}

	return "", "", false
}

func getandparsefile(path string) (urlList []string, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// windows new line problem solved.
	fixedURLList := strings.ReplaceAll(string(data), "\r\n", "\n")
	urlList = strings.Split(fixedURLList, "\n")
	return urlList, nil
}

func main() {
	arguments := os.Args[1:]
	url, path, status := operatecommand(arguments)

	if status != true {
		fmt.Println("An unexpected error occurred.")
		return
	}

	if len(url) > 0 {
		getstatuscode(url)
	} else if len(path) > 0 {
		if urlList, err := getandparsefile(path); err == nil {
			for _, url := range urlList {
				getstatuscode(url)
			}
		}
	}

}
