package main_test

import (
	"testing"

	cli "github.com/freddy50806/gRPCTestWithGomock/Unary/client"
	exmock "github.com/freddy50806/gRPCTestWithGomock/Unary/mock_example"
	pb "github.com/freddy50806/gRPCTestWithGomock/Unary/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSendZero(t *testing.T) {
	//set controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//set mock client stub
	mockClient := exmock.NewMockExampleServiceClient(ctrl)
	mockClient.EXPECT().IsEven(gomock.Any(), &pb.Request{Num: 0}).Return(&pb.Reply{Answer: true}, nil)
	//test
	testSendZero(t, mockClient)
}

func testSendZero(t *testing.T, client pb.ExampleServiceClient) {
	answer, err := cli.SendZero(client)
	assert.EqualValues(t, nil, err)
	assert.Equal(t, true, answer, "Zero should be even!")
}
