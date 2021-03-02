package main

import (
	"context"
	"fmt"
	"github.com/niels-baltodano/grpc-pb/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + " " + lastname

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil

}

func main() {
	fmt.Println("Hello Server Greet")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

}
