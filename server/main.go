package main

import (
	"context"
	pb "example/grpcdemo/proto"
	"fmt"
	"io"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	fmt.Println("The new client is connected")
	return &pb.HelloResponse{
		Name: "Aravind",
	}, nil
}

func (s *helloServer) SayBye(req *pb.NoParam, stream pb.GreetService_SayByeServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&pb.HelloResponse{
			Name: "This is the world",
		})
		time.Sleep(time.Second)
	}
	return nil
}
func (s *helloServer) SaystreamHello(stream pb.GreetService_SaystreamHelloServer) error {
	for {
		message, err := stream.Recv()
		if err == io.EOF {

			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("The message Recived from Client:", message.Name)
		stream.Send(&pb.HelloResponse{
			Name: "Aravind likes ramya",
		})
	}

	return nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	fmt.Println("The server is listening on port [:8080]")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	fmt.Println("The plain data:")
}
