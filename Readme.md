Mocking Service Example for gRPC
===

It's use [this documentation](https://github.com/grpc/grpc-go/blob/master/Documentation/gomock-example.md) for reference, and simulate three different situation to test.

Install
---
Before you install, be sure you have setuped Golang 1.6 or higher on your computer.

```
$go get -u google.golang.org/grpc
$go get -u github.com/golang/protobuf/protoc-gen-go

$go get github.com/golang/mock/gomock
$go install github.com/golang/mock/mockgen
```

Run Test
---
Change directory to the file you want, generate grpc code and gomcok code


#### Unary

```
$cd $GOPATH/src/github.com/freddy50806/gRPCTestWithGomock/Unary
$protoc -I proto proto/example.proto --go_out=plugins=grpc:proto
```
Use gomock gernerate mock interface in the `proto/` which grpc gernerate.
```
$mockgen github.com/freddy50806/gRPCTestWithGomock/Unary/proto ExampleServiceClient > mock_example/mock_example.go
```
Run test in `client/main_test.go`

#### Streaming
```
$cd $GOPATH/src/github.com/freddy50806/gRPCTestWithGomock/Streaming
$protoc -I proto proto/example.proto --go_out=plugins=grpc:proto
```
Use gomock gernerate mock interface in the `proto/` which grpc gernerate.
```
$mockgen github.com/freddy50806/gRPCTestWithGomock/Streaming/proto ExampleServiceClient,ExampleService_DisplayLinesClient > mock_example/mock_example.go
```
Run test in `client/main_test.go`

#### Bi-Streaming
```
$cd $GOPATH/src/github.com/freddy50806/gRPCTestWithGomock/Bi-Streaming
$protoc -I proto proto/example.proto --go_out=plugins=grpc:proto
```
Use gomock gernerate mock interface in the `proto/` which grpc gernerate.
```
$mockgen github.com/freddy50806/gRPCTestWithGomock/Bi-Streaming/proto ExampleServiceClient,ExampleService_StreamingIsEvenClient > mock_example/mock_example.go
```
Run test in `client/main_test.go`