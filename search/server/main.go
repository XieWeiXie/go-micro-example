package main

import (
	"context"
	"log"
	"net"

	"github.com/wuxiaoxiaoshen/go-micro-example/search/protoc"
	"google.golang.org/grpc"
)

type Server struct {
}

func NewItems() []search.Response {
	return []search.Response{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
	}
}

func (s *Server) FindById(cx context.Context, in *search.Request) (out *search.Response, err error) {
	item := NewItems()[0]
	return &item, nil
}

func main() {
	port := ":1234"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ss := Server{}
	search.RegisterGreeterServer(s, &ss)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
