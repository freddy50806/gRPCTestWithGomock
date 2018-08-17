package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/freddy50806/gRPCTestWithGomock/Unary/proto"
	"google.golang.org/grpc"
)

var addr = flag.String("server_addr", "127.0.0.1:555", "The server address in the format of host:port")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewExampleServiceClient(conn)
	SendZero(client)
}
func SendZero(client pb.ExampleServiceClient) (bool, error) {
	r, err := client.IsEven(context.Background(), &pb.Request{Num: 0})
	if err != nil {
		return false, err
	}
	fmt.Println(r.Answer)
	if r.Answer {
		return true, nil
	}
	return false, nil
}
