package main

import (
	"log"
	"time"
)

func salutation() string {
	return "Ciao mondo!"
}

func main() {
	for {
		log.Printf(salutation())
		time.Sleep(1 * time.Second)
	}
}
