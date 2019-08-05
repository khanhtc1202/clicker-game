package grpc

import (
	game "clickers/server/api/v1"
	"clickers/server/s-game-engine/internal/server/logic"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *GRpc {
	return &GRpc{
		address: address,
	}
}

func (g *GRpc) GetSize(ctx context.Context, req *game.GetSizeRequest) (*game.GetSizeResponse, error) {
	log.Println("GetSize in s-game-engine is called")
	return &game.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (g *GRpc) SetScore(ctx context.Context, req *game.SetScoreRequest) (*game.SetScoreResponse, error) {
	log.Println("SetScore in s-game-engine is called")
	return &game.SetScoreResponse{
		Set: logic.SetScore(req.Score),
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
	game.RegisterGameEngineServer(g.srv, g)
	log.Println("address:", g.address)
	log.Println("Starting gRPC server for s-game-engine")

	err = g.srv.Serve(lis)
	if err != nil {
		log.Println("Error: Failed to serve s-game-engine service")
	}

	return nil
}
