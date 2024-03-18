package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: make my root directory configurable
		filename := "./content" + r.URL.Path

		// add default document
		if filename == "./content/" {
			filename = "./content/index.html"
		}

		// TODO: maybe improve the logging a bit ;)
		fmt.Println(filename)

		body, _ := os.ReadFile(filename)

		if strings.HasSuffix(filename, ".css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(filename, ".svg") {
			w.Header().Set("Content-Type", "image/svg+xml")
		} else if strings.HasSuffix(filename, ".html") {
			w.Header().Set("Content-Type", "text/html")

			// convert to string and do some basic SSI
			bodyString := string(body)

			idx := strings.Index(bodyString, "<!--#include file=")
			for idx != -1 {
				idx2 := strings.Index(bodyString, "-->")
				subfile := bodyString[idx+19 : idx2-1]

				subfileContent, _ := os.ReadFile("./content" + subfile)

				newBodyString := bodyString[0:idx] + string(subfileContent) + bodyString[idx2+3:len(bodyString)]
				bodyString = newBodyString
				idx = strings.Index(bodyString, "<!--#include file=")
			}

			body = []byte(bodyString)
		} else if strings.HasSuffix(filename, ".js") {
			w.Header().Set("Content-Type", "text/javascript")
		}

		w.Write(body)
	})

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
