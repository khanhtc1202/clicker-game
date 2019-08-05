package main

import (
	game "clickers/server/api/v1"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var addressStr = flag.String("address", "localhost:50051", "server address to request")
	flag.Parse()

	conn, err := grpc.Dial(*addressStr, grpc.WithInsecure())
	if err != nil {
		log.Println("address:", *addressStr)
		log.Println(err)
		log.Fatal("Error: can not connect to server")
	}

	cli := game.NewGameClient(conn)
	if cli == nil {
		log.Println("Client is nil")
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
			log.Fatal("Error: Failed to close the connection")
		}
	}()

	var req game.GetHighScoreRequest
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := cli.GetHighScore(timeoutCtx, &req)
	if err != nil {
		log.Println("Error: Can not get response")
		log.Fatal(err)
	}

	if res != nil {
		log.Println("HighScore:", res.GetHighScore())
	} else {
		log.Println("Could not get highscore")
	}
}
