package main

import (
	"log"

	"github.com/sibelephant/go-grpc-demo/internal/db"
	"github.com/sibelephant/go-grpc-demo/internal/rocket"
	grpc "github.com/sibelephant/go-grpc-demo/internal/transport"
)

func Run() error {
	//responsible for initializing and starting our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("failed to run migrations")
		return err
	}

	rktService := rocket.New(rocketStore)

	rktHandler := grpc.New(rktService)
	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}
func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
