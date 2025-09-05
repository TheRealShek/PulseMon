package main

import (
	"fmt"
	"log"
	"main/configs"
)

func main() {
	cfg, err := configs.LoadConfig("../configs/pulsemon.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", cfg)
}
