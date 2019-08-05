package main

import (
	"clickers/server/s-bff/bff"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	grpcAddressHighScore := flag.String(
		"address-m-highscore", "localhost:50051", "the grpc server for highscore service")
	grpcAddressGameEngine := flag.String(
		"address-m-game-engine", "localhost:60051", "the grpc server for game-engine service")
	serverAddress := flag.String("addess-http", ":8081", "HTTP server address")

	flag.Parse()

	gameCli, err := bff.NewGrpcGameServiceClient(*grpcAddressHighScore)
	if err != nil {
		log.Fatal("Error: can not connect to s-highscore, got:", err)
	}
	gameEngineCli, err := bff.NewGrpcGameEngineServiceClient(*grpcAddressGameEngine)
	if err != nil {
		log.Fatal("Error: can not connect to s-game-engine, got:", err)
	}

	gr := bff.NewGameResource(gameCli, gameEngineCli)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
	}))

	router.GET("/hs", gr.GetHighScore)
	router.POST("/hs/:hs", gr.SetHighScore)
	router.GET("/size", gr.GetSize)
	router.POST("/score/:score", gr.SetScore)

	err = router.Run(*serverAddress)
	if err != nil {
		log.Fatal("Error: Can not start http server, got:", err)
	}

	log.Println("Started http server at", *serverAddress)
}
