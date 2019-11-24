package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

func printStatusCodeWithColor(code int) {
	white := color.New(color.FgWhite)
	black := color.New(color.FgBlack)

	if code >= 200 && code <= 226 {
		white.Add(color.BgGreen).Printf(" %d ", code)
	} else if code >= 300 && code <= 308 {
		black.Add(color.BgYellow).Printf(" %d ", code)
	} else if code >= 400 && code <= 499 {
		white.Add(color.BgRed).Printf(" %d ", code)
	} else {
		black.Add(color.BgWhite).Printf(" %d ", code)
	}

}

func getStatusCode(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		printStatusCodeWithColor(resp.StatusCode)
		fmt.Println(" ", url)
	}
	resp.Body.Close()
}

func operateCommand(args []string) (url, path string, status bool) {
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

func getAndParseFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// windows new line problem solved.
	fixedURLList := strings.ReplaceAll(string(data), "\r\n", "\n")
	urlList := strings.Split(fixedURLList, "\n")
	return urlList, nil
}

func main() {
	arguments := os.Args[1:]
	url, path, status := operateCommand(arguments)

	if status != true {
		fmt.Println("An unexpected error occurred.")
		return
	}

	if len(url) > 0 {
		getStatusCode(url)
	} else if len(path) > 0 {
		if urlList, err := getAndParseFile(path); err == nil {
			for _, url := range urlList {
				getStatusCode(url)
			}
		}
	}

}
