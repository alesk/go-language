package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		fmt.Printf("Fetching %v\n", url)

		// Ctrl+Shift I for quick documentation
		if !strings.HasPrefix(url,"http://")  {
			url = "http://" + url
		}


		response, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching \033[35m%v\033[0m: %v\n", url, err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading body %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n\nStatus code: %d\n", response.StatusCode)
	}

}
