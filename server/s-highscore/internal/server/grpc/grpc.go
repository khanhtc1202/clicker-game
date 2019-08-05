package grpc

import (
	game "clickers/server/api/v1"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var highScore = 9999999999.0

type GRpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *GRpc {
	return &GRpc{
		address: address,
	}
}

func (g *GRpc) SetHighScore(ctx context.Context, input *game.SetHighScoreRequest) (*game.SetHighScoreResponse, error) {
	log.Println("SetHighScore in s-highscore is called")
	highScore = input.HighScore
	return &game.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *GRpc) GetHighScore(ctx context.Context, input *game.GetHighScoreRequest) (*game.GetHighScoreResponse, error) {
	log.Println("GetHighScore in s-highscore is called")
	return &game.GetHighScoreResponse{
		HighScore: highScore,
	}, nil
}

func (g *GRpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		log.Println("Error: Failed to open tcp port")
		return err
	}

	srvOpts := make([]grpc.ServerOption, 0)

	g.srv = grpc.NewServer(srvOpts...)
	game.RegisterGameServer(g.srv, g)
	log.Println("address:", g.address)
	log.Println("Starting gRPC server for s-highscore")

	err = g.srv.Serve(lis)
	if err != nil {
		log.Println("Error: Failed to serve s-highscore service")
	}

	return nil
}
