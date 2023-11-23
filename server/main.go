package main

import (
	"context"
	helloworld "golang-grpc/protobuf/helloworld/protobuf"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption

	helloWorldServer := &HelloWorldServer{}

	grpcServer := grpc.NewServer(opts...)
	helloworld.RegisterHelloWorldServer(grpcServer, helloWorldServer)

	grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

type HelloWorldServer struct {
	helloworld.UnimplementedHelloWorldServer
}

func (h *HelloWorldServer) SayHello(ctx context.Context, request *helloworld.Request) (*helloworld.Response, error) {
	return &helloworld.Response{Message: "message received: " + request.Message}, nil
}

func (h *HelloWorldServer) ChatSayHello(stream helloworld.HelloWorld_ChatSayHelloServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Print(err)
			return err
		}

		err = stream.Send(&helloworld.Response{
			Message: "message received: " + request.Message,
		})

		if err != nil {
			log.Print(err)
			return err
		}

		go newFunction(stream)
	}
}

func newFunction(stream helloworld.HelloWorld_ChatSayHelloServer) {
	for i := 0; i < 2; i++ {
		time.Sleep(3 * time.Second)

		err := stream.Send(&helloworld.Response{
			Message: "the actual time is: " + time.Now().String(),
		})

		if err != nil {
			log.Print(err)
		}
	}
}
