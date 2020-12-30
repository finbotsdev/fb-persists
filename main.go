package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("/data/fp-persists.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println("more text to append")

	for {
		logger.Println("fb-persists")
		fmt.Println("fb-persists")

		time.Sleep(time.Second * 2)
	}
}
