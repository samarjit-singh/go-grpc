package main

import (
	pb "go-grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const(
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// defer ensure that this function runs at the very end of execution
	 defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// name := &pb.NamesList{
	// 	Names: []string{"Samar", "Alice", "Bob"},
	// }

	callSayHello(client)
}