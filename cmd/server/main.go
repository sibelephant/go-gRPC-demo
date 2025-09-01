package main

import (
	"log"
	"time"

	"github.com/sibelephant/go-grpc-demo/internal/db"
	"github.com/sibelephant/go-grpc-demo/internal/rocket"
	grpc "github.com/sibelephant/go-grpc-demo/internal/transport"
)

func Run() error {
	//responsible for initializing and starting our gRPC server
	var rocketStore db.Store
	var err error

	// Retry database connection up to 30 times (30 seconds)
	for i := 0; i < 30; i++ {
		rocketStore, err = db.New()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/30): %v", i+1, err)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to database after 30 attempts")
		return err
	}

	log.Println("Successfully connected to database")

	// Temporarily disable migrations - table already exists
	// err = rocketStore.Migrate()
	// if err != nil {
	// 	log.Println("failed to run migrations")
	// 	return err
	// }

	rktService := rocket.New(rocketStore)

	rktHandler := grpc.New(rktService)
	log.Println("Starting gRPC server on port 50051...")
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
