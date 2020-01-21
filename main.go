package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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
	url := flag.String("url", "", "Give an url")
	path := flag.String("path", "", "Give a path")

	flag.Parse()

	if len(*url) > 0 {
		getStatusCode(*url)
	} else if len(*path) > 0 {
		if urlList, err := getAndParseFile(*path); err == nil {
			for _, url := range urlList {
				getStatusCode(url)
			}
		}
	} else {
		fmt.Println("You must specify a file path or url.")
	}

}
