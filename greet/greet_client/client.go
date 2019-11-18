package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hugolesta/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client %f", c)

	// doUnary(c)

	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hugo",
			LastName:  "Lesta",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error when calling greet RPC: %v", err)
	}
	log.Printf("Response from greet %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server stream RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hugo",
			LastName:  "Lesta",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			// We're reach end of the stream server

			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}

		log.Printf("Response GreetManyTimes: %v", msg.GetResult())
	}

}
