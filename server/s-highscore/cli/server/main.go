package main

import (
	"clickers/server/s-highscore/internal/server/grpc"
	"flag"
	"log"
)

func main() {
	var addressStr = flag.String("address", ":50051", "address serve s-highscore service")
	flag.Parse()

	server := grpc.NewServer(*addressStr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
