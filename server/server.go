package server

import (
	"context"
	"github.com/Wilddogmoto/grpcTest/sender"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
)

type GRPCServer struct{}

func InitServer() {

	log.Info("GRPC server started")

	listener, err := net.Listen("tcp", ":53003")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	sender.RegisterSenderServer(grpcServer, &GRPCServer{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		grpclog.Fatalf("failed to server: %v", err)
	}
}

func (s *GRPCServer) SendMessage(ctx context.Context, in *sender.InMessage) (out *sender.OutMessage, err error) {
	log.Infof("income message: %s", in.Message)

	out = &sender.OutMessage{}

	out.Message = "3324636367533"

	//	out = &sender.OutMessage{
	//		Message: "3324636367533",
	//	}

	return
}
