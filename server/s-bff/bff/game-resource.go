package bff

import (
	game "clickers/server/api/v1"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

type gameResouse struct {
	gameClient       game.GameClient
	gameEngineClient game.GameEngineClient
}

func NewGameResource(
	gameClient game.GameClient,
	gameEngineClient game.GameEngineClient,
) *gameResouse {
	return &gameResouse{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}

func NewGrpcGameServiceClient(srvAddress string) (game.GameClient, error) {
	conn, err := grpc.Dial(srvAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		log.Println("Connected to s-highscore service", srvAddress)
	}

	client := game.NewGameClient(conn)
	return client, nil
}

func NewGrpcGameEngineServiceClient(srvAddress string) (game.GameEngineClient, error) {
	conn, err := grpc.Dial(srvAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		log.Println("Connected to s-game-engine service", srvAddress)
	}

	client := game.NewGameEngineClient(conn)
	return client, nil
}

func (gr *gameResouse) SetHighScore(c *gin.Context) {
	highScoreString := c.Param("hs")
	highScoreFloat64, err := strconv.ParseFloat(highScoreString, 64)
	if err != nil {
		log.Println("Error: Failed to convert high score, value:", highScoreString)
		return
	}
	_, err = gr.gameClient.SetHighScore(context.Background(), &game.SetHighScoreRequest{
		HighScore: highScoreFloat64,
	})
	if err != nil {
		log.Println("Error: can not set high score to s-highscore")
		log.Println(err)
	}
}

func (gr *gameResouse) GetHighScore(c *gin.Context) {
	res, err := gr.gameClient.GetHighScore(context.Background(), &game.GetHighScoreRequest{})
	if err != nil {
		log.Println("Error: error while getting high score")
		log.Println(err)
		return
	}
	hsString := strconv.FormatFloat(res.GetHighScore(), 'e', -1, 64)
	c.JSONP(http.StatusOK, gin.H{
		"hs": hsString,
	})
}

func (gr *gameResouse) GetSize(c *gin.Context) {
	res, err := gr.gameEngineClient.GetSize(context.Background(), &game.GetSizeRequest{})
	if err != nil {
		log.Println("Error: error while getting size")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"size": res.GetSize(),
	})
}

func (gr *gameResouse) SetScore(c *gin.Context) {
	scoreString := c.Param("score")
	scoreFloat64, err := strconv.ParseFloat(scoreString, 64)
	if err != nil {
		log.Println("Error: Failed to convert score, value:", scoreString)
		return
	}

	_, err = gr.gameEngineClient.SetScore(context.Background(), &game.SetScoreRequest{
		Score: scoreFloat64,
	})
	if err != nil {
		log.Println("Error: can not set score to s-game-engine")
		log.Println(err)
	}
}
