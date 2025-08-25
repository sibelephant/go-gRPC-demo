//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/sibelephant/go-grpc-demo/internal/rocket Store


package rocket

import (
	"context"
)

// rocket should contain the definition of our rocket
type Rocket struct {
	ID     string
	Name   string
	Type   string
	Flight int
}

// Store defines the interface we expect
// our db implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) (Rocket, error)
}

// service responsible for updating the rocket inventory
type Service struct {
	Store Store
}

// New returns a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID retrives a rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket	inserts a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

func (s Service) DeleteRocket(ctx context.Context, id string) error {
	_, err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}