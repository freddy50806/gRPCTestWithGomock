package main_test

import (
	"io"
	"testing"

	cli "github.com/freddy50806/gRPCTestWithGomock/Streaming/client"
	exmock "github.com/freddy50806/gRPCTestWithGomock/Streaming/mock_example"
	pb "github.com/freddy50806/gRPCTestWithGomock/Streaming/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	//set controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//set mock client stub
	mockStream := exmock.NewMockExampleService_DisplayLinesClient(ctrl)
	gomock.InOrder(
		mockStream.EXPECT().Recv().Return(&pb.Reply{Data: "First	Line"}, nil),
		mockStream.EXPECT().Recv().Return(&pb.Reply{Data: "Second	Line"}, nil),
		mockStream.EXPECT().Recv().Return(&pb.Reply{Data: "Third	Line"}, nil),
		mockStream.EXPECT().Recv().Return(&pb.Reply{Data: "Fourth	Line"}, nil),
		mockStream.EXPECT().Recv().Return(nil, io.EOF),
	)
	mockClient := exmock.NewMockExampleServiceClient(ctrl)
	mockClient.EXPECT().DisplayLines(gomock.Any(), gomock.Any()).Return(mockStream, nil)
	//test
	testDisplay(t, mockClient)
}
func testDisplay(t *testing.T, client pb.ExampleServiceClient) {
	answer, err := cli.Display(4, client)
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, 4, answer)
}
