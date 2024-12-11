package main

import (
	"context"
	"log"
	"net"

	pb "Animales-go/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAnimalesserviceServer
}

func (s *server) GetAnimalesInfo(ctx context.Context, req *pb.AnimalRequest) (*pb.AnimalResponse, error) {
	animaldata := map[string]pb.AnimalResponse{
		"Perro": {Name: "Perro", Type: "canino", NumPatas: 4},
		"Gato":  {Name: "Gato", Type: "felino", NumPatas: 4},
	}

	if res, ok := animaldata[req.Name]; ok {
		return &res, nil
	}

	return &pb.AnimalResponse{Name: "Unknown", Type: "Unknown", NumPatas: 0}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAnimalesserviceServer(grpcServer, &server{})

	log.Println("Starting server on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}
