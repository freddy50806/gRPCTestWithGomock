package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	pb "github.com/freddy50806/gRPCTestWithGomock/Bi-Streaming/proto"
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
	SendASeriesOfNumber(4, client)
}
func SendASeriesOfNumber(num int32, client pb.ExampleServiceClient) (int, error) {
	stream, err := client.StreamingIsEven(context.Background())
	if err != nil {
		log.Fatalf("err:%v", err)
	}
	var count int32
	waitc := make(chan struct{})
	//Run Recieve
	go func() {
		fmt.Println("Start to recieve!")
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				fmt.Println("Recieve close!")
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Recieve Error : %v", err)
			}
			fmt.Printf("Got message %t\n", in.Answer)
		}
	}()

	for count = 0; count <= num; count++ {
		fmt.Println("Send", count)
		if err := stream.Send(&pb.Request{Num: count}); err != nil {
			return int(count), err
		}
	}

	stream.CloseSend()
	<-waitc
	fmt.Println("Close!!")
	return int(count), nil
}
