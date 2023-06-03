package main

import (
	"conduit-go/config"
	"conduit-go/internal/app"
	"fmt"
	"log"
)

func main() {
	fmt.Println("start!")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("failed to read config: ", err)
	}
	app.Run(cfg)
}
