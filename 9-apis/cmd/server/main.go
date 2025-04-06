package main

import (
	"log"

	"github.com/felipeazsantos/pos-goexpert/apis/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
}