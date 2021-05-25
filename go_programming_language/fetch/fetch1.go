package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FixPrefix(url string, prefix string) string {
	if !strings.HasPrefix(url, prefix) {
		url = prefix + url
	}
	return url
}

func main() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		url := FixPrefix(url, prefix)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading")
			os.Exit(1)
		}
		fmt.Print("\n\n")
		fmt.Printf("Status code: %s", resp.Status)
		fmt.Printf("%s", b)
	}
}