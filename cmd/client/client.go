package main

import (
	"context"
	protoBuffer "grpc-haversine/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not to connect to gRPC server %v", err)
	}

	client := protoBuffer.NewHaversineServiceClient(connection)

	defer connection.Close()

	request := &protoBuffer.HaversineRequest{
		FirstCoords: &protoBuffer.Coords{
			Lat:  -19.8157,
			Long: -43.9542,
		},

		SecondCoords: &protoBuffer.Coords{
			Lat:  -23.5489,
			Long: -46.6388,
		},
	}

	response, err := client.Haversine(context.Background(), request)
	if err != nil {
		log.Fatalf("error during the request %v", err)
	}

	log.Printf("%s %.3f %s", "approximately", response.GetDistance()/1000, "km")
}
