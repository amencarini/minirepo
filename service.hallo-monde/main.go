package main

import (
	"log"
	"time"
)

func salutation() string {
	return "Hallo monde!"
}

func main() {
	for {
		log.Printf(salutation())
		time.Sleep(1 * time.Second)
	}
}
