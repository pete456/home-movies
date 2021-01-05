package main

import (
	"home-movies/internal/net"
	"home-movies/config"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(config.Logpath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	net.InitServer()
}
