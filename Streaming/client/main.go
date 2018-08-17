package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	pb "github.com/freddy50806/gRPCTestWithGomock/Streaming/proto"
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
	Display(32, client)
}
func Display(num int32, client pb.ExampleServiceClient) (int, error) {
	count := 0
	stream, err := client.DisplayLines(context.Background(), &pb.Request{Num: num})
	if err != nil {
		log.Fatalf("err:%v", err)
	}
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return count, err
		}
		count++
		fmt.Println("Line ", count, " : ", l.Data)
	}
}
