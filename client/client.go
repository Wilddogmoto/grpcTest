package client

import (
	"context"
	"fmt"
	"github.com/Wilddogmoto/grpcTest/sender"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func ClientGRPC() {

	connect, err := grpc.Dial("127.0.0.1:53003", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	client := sender.NewSenderClient(connect)
	request := &sender.InMessage{Message: "123"}

	response, err := client.SendMessage(context.Background(), request)
	if err != nil {

		grpclog.Fatalf("fail to send message: %v", err)
	}

	fmt.Print(response.GetMessage())
}
