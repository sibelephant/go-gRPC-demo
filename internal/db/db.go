package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sibelephant/go-grpc-demo/internal/rocket"
)

type Store struct {
	db *sqlx.DB
}

// New returns a new store or eror
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUsername, dbTable, dbPassword, dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}

	return Store{
		db: db,
	}, nil
}

// GetRocketById retrives a rocket by id
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	err := s.db.QueryRow(
		`SELECT id, name, type, flight FROM rockets WHERE id=$1`,
		id,
	).Scan(&rkt.ID, &rkt.Name, &rkt.Type, &rkt.Flight)

	if err != nil {
		return rocket.Rocket{}, err
	}
	return rkt, nil
}

// this inserts a rocket into a rocket table
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	var insertedRocket rocket.Rocket
	err := s.db.QueryRow(
		`INSERT INTO rockets (name, type, flight) 
		 VALUES($1, $2, $3) 
		 RETURNING id, name, type, flight`,
		rkt.Name, rkt.Type, rkt.Flight,
	).Scan(&insertedRocket.ID, &insertedRocket.Name, &insertedRocket.Type, &insertedRocket.Flight)

	if err != nil {
		return rocket.Rocket{}, errors.New("failed to insert into database")
	}

	return insertedRocket, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}
