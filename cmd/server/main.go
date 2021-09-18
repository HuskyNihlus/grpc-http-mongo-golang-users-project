package main

import (
	"PROJECT/pkg/adder"
	"PROJECT/pkg/api"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	fmt.Println("Mongo Connect")

	err = client.Connect(context.TODO())
	handleError(err)
	collection := client.Database("usersdb").Collection("users")

	s := grpc.NewServer()
	srv := adder.New(collection)
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", "localhost:8080")
	handleError(err)
	fmt.Println("Server ON")

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}