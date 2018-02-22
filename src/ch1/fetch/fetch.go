package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		fmt.Printf("Fetching %v\n", url)

		response, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching \033[35m%v\033[0m: %v\n", url, err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading body %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s", body)
	}

}
