package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
	"strings"
)

/*
To fetch a long filename list, use:

	curl https://moz.com/top500/domains/csv -o topsites.csv
	cat topsites.csv | awk -F "," '{print $2}' | xargs ./fetchall

*/

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Erro fetching %s: %v\n", url, err)
		return
	}

	// check if file exists
	slug := slugify(url)
	file_index := 1
	filename := func() string { return fmt.Sprintf("tmp/%s_%03d.html", slug, file_index) }
	for {
		_, err := os.Stat(filename())
		if os.IsNotExist(err) {
			break
		}
		file_index += 1
	}

	f, err := os.Create(filename())
	defer f.Close()
	if err != nil {
		ch <- fmt.Sprintf("Erro writng content for %s: %v\n", url, err)
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func slugify(s string) string {
	r := regexp.MustCompile("[^A-Za-z0-9]+")
	return r.ReplaceAllString(s, "_")
}
