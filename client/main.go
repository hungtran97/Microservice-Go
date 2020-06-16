package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "grpc/pb/user"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	request := &pb.GetRequest{
		Keyword: "test",
	}
	//fmt.Println(context.Background())

	resp, err := client.GetUsers(context.Background(), request)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Users[0].Name)
}
