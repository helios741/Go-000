package main

import (
	pb "Week04/api/students/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)
const (
	port = ":50052"
)
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStudentServer(s, initServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
