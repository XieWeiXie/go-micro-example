package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wuxiaoxiaoshen/go-micro-example/search/protoc"

	"google.golang.org/grpc"
)

const (
	address = "localhost:1234"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := search.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var request search.Request
	request.Id = "1"
	response, err := c.FindById(ctx, &request)
	fmt.Println(response, err)
}
