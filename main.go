package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	queryString := strings.Join(os.Args[1:], "&")
	queries, err := url.ParseQuery(queryString)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for key, value := range queries {
		rawKey, err := url.PathUnescape(key)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// NOTE: a=1&a=2&a=3
		rawValue, err := url.PathUnescape(strings.Join(value, ","))
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Printf("%s: %s\n", rawKey, rawValue)
	}
}
