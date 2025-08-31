package grpc

import (
	"context"
	"log"
	"net"

	"github.com/sibelephant/go-grpc-demo/internal/rocket"
	rkt "github.com/sibelephant/go-grpc-demo/rocket/v1"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
	rkt.UnimplementedRocketServiceServer
}

// New returns a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil
}

// GetRocket handles GetRocket gRPC requests
func (h Handler) GetRocket(ctx context.Context, r *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	rocket, err := h.RocketService.GetRocketByID(ctx, r.Id)
	if err != nil {
		return &rkt.GetRocketResponse{}, err
	}

	protoRocket := &rkt.Rocket{
		Id:   rocket.ID,
		Name: rocket.Name,
		Type: rocket.Type,
	}

	return &rkt.GetRocketResponse{
		Rocket: protoRocket,
	}, nil
}

// AddRocket handles AddRocket gRPC requests
func (h Handler) AddRocket(ctx context.Context, r *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	newRocket := rocket.Rocket{
		ID:   r.Rocket.Id,
		Name: r.Rocket.Name,
		Type: r.Rocket.Type,
	}

	rocket, err := h.RocketService.InsertRocket(ctx, newRocket)
	if err != nil {
		return &rkt.AddRocketResponse{}, err
	}

	protoRocket := &rkt.Rocket{
		Id:   rocket.ID,
		Name: rocket.Name,
		Type: rocket.Type,
	}

	return &rkt.AddRocketResponse{
		Rocket: protoRocket,
	}, nil
}

// DeleteRocket handles DeleteRocket gRPC requests
func (h Handler) DeleteRocket(ctx context.Context, r *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	err := h.RocketService.DeleteRocket(ctx, r.Rocket.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{
			Status: "failed to delete rocket",
		}, err
	}

	return &rkt.DeleteRocketResponse{
		Status: "successfully deleted rocket",
	}, nil
}
