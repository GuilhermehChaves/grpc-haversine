package main

import (
	protoBuffer "grpc-haversine/pb"
	"grpc-haversine/server"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	protoBuffer.RegisterHaversineServiceServer(grpcServer, &server.Positions{})

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Cannot start server %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Cannot start server %v", err)
	}
}
