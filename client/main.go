package main

import (
	"context"
	pb "example/grpcdemo/proto"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func clientSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Name)
}

func clientSayBye(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()

	stream, _ := client.SayBye(ctx, &pb.NoParam{})
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
	fmt.Println("The server stream ended")

}

func clientSendStream(client pb.GreetServiceClient) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	stream, _ := client.SaystreamHello(context.Background())
	for i := 0; i < 5; i++ {
		stream.Send(&pb.HelloRequest{
			Name: "My baby Ramya",
		})
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
	time.Sleep(time.Second)
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
	fmt.Println("Bidirectional Streaming Done")
}
func main() {
	lis, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	client := pb.NewGreetServiceClient(lis)

	//clientSayHello(client)
	//clientSayBye(client)
	clientSendStream(client)
}
