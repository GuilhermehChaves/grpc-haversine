package server

import (
	"context"
	protoBuffer "grpc-haversine/pb"
	"math"
)

type Positions struct{}

func (p *Positions) Haversine(ctx context.Context, in *protoBuffer.HaversineRequest) (*protoBuffer.HaversineResponse, error) {
	distance := calculate(in)
	return &protoBuffer.HaversineResponse{
		Distance: distance,
	}, nil
}

func calculate(locations *protoBuffer.HaversineRequest) float64 {
	//To Radians
	locations.FirstCoords.Lat = locations.FirstCoords.Lat * math.Pi / 180
	locations.FirstCoords.Long = locations.FirstCoords.Long * math.Pi / 180
	locations.SecondCoords.Lat = locations.SecondCoords.Lat * math.Pi / 180
	locations.SecondCoords.Long = locations.SecondCoords.Long * math.Pi / 180

	//Haversine Formula
	r := 6371e3
	deltaLat := locations.SecondCoords.GetLat() - locations.FirstCoords.GetLat()
	deltaLong := locations.SecondCoords.GetLong() - locations.FirstCoords.GetLong()

	a := math.Pow(math.Sin(deltaLat/2), 2) +
		math.Cos(locations.FirstCoords.GetLat())*
			math.Cos(locations.SecondCoords.GetLat())*
			math.Pow(math.Sin(deltaLong/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return r * c
}
