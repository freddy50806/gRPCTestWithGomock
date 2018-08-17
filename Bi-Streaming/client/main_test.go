package main_test

import (
	"io"
	"testing"

	cli "github.com/freddy50806/gRPCTestWithGomock/Bi-Streaming/client"
	exmock "github.com/freddy50806/gRPCTestWithGomock/Bi-Streaming/mock_example"
	pb "github.com/freddy50806/gRPCTestWithGomock/Bi-Streaming/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	//set controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//set mock client stub
	mockStream := exmock.NewMockExampleService_StreamingIsEvenClient(ctrl)

	mockClient := exmock.NewMockExampleServiceClient(ctrl)
	// mockStream.EXPECT().Recv().Return(nil, fmt.Errorf("isn't not true"))
	// first
	firstCall := mockStream.EXPECT().Send(&pb.Request{Num: 0})
	// second
	secondCall := mockStream.EXPECT().Send(&pb.Request{Num: 1})
	// close
	closeCall := mockStream.EXPECT().CloseSend().Return(nil)
	// recieve call sequential
	gomock.InOrder(
		mockStream.EXPECT().Recv().Return(&pb.Reply{Answer: true}, nil).After(firstCall),
		mockStream.EXPECT().Recv().Return(&pb.Reply{Answer: false}, nil).After(secondCall),
		mockStream.EXPECT().Recv().Return(nil, io.EOF).After(closeCall),
	)
	mockClient.EXPECT().StreamingIsEven(gomock.Any(), gomock.Any()).Return(mockStream, nil)

	//test
	testStreamingIsEven(t, mockClient)
	return
}
func testStreamingIsEven(t *testing.T, client pb.ExampleServiceClient) {
	answer, err := cli.SendASeriesOfNumber(1, client)
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, 2, answer)
	return
}
