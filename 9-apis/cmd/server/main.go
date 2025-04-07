package main

import (
	"fmt"
	"log"

	"github.com/felipeazsantos/pos-goexpert/apis/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.DBDriver)
}