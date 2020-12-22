package services

import (
	"context"
	"github.com/jackc/pgx/pgxpool"
	"log"
	"queue/databases/postgres"
	"queue/models"
)

type MaintenanceService struct {
	pool *pgxpool.Pool
}

func NewMaintenanceService(pool *pgxpool.Pool) *MaintenanceService {
	return &MaintenanceService{pool: pool}
}

func (receiver *MaintenanceService) AddCity(City models.City) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Cant connect")
		return
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddCity, City.ID, City.Name)
	if err != nil {
		log.Fatal("cant add city")
	}
	return nil
}
func (receiver *MaintenanceService) AddBranch(Branch models.Branch) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Cant connect")
		return
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddBranch, Branch.ID, Branch.Address, Branch.CityID)
	if err != nil {
		log.Fatal("cant add branch")
	}
	return nil
}