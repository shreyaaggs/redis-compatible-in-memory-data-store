package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/shreyaaggs/redis-compatible-in-memory-data-store/internal/server"
)

func main() {
	fmt.Println("redis-go starting")
	port := flag.Int("port", 6380, "port to listen on")
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	srv := server.New(addr)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
