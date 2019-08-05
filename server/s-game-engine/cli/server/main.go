package main

import (
	"clickers/server/s-game-engine/internal/server/grpc"
	"flag"
	"log"
)

func main() {
	var addressStr = flag.String("address", ":60051", "address serve s-game-engine service")
	flag.Parse()

	server := grpc.NewServer(*addressStr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
