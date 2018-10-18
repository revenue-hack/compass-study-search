package main

import (
	"flag"
	"log"
	"strings"
	"os"
)


func main() {
	var inputWords string
	flag.StringVar(&inputWords, "keywords", "", "input keywords e.g go,python")
	flag.Parse()

	if inputWords == "" {
		log.Println("input keywords e.g go,python")
		os.Exit(1)
	}

	keywords := strings.Split(inputWords, ",")
	if len(keywords) == 0 {
		log.Println("input keywords e.g go,python")
		os.Exit(1)
	}

}
