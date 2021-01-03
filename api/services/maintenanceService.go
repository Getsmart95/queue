package services

import "C"
import (
	"github.com/jackc/pgx/pgxpool"
	"queue/databases/postgres"
	"queue/models"
	"queue/tokens"
	"context"
	"log"
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
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddCity, City.Name)
	if err != nil {
		return
	}
	return nil
}

func (receiver *MaintenanceService) GetAllCities()(Cities []models.City, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetAllCities)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next(){
		City := models.City{}
		errRole := rows.Scan(&City.ID, &City.Name)
		if errRole != nil {
			return
		}
		Cities = append(Cities, City)
	}
	return Cities, nil
}

func (receiver *MaintenanceService) AddBranch(Branch models.Branch) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddBranch, Branch.Address, Branch.CityID)
	if err != nil {
		return
	}
	return nil
}

func (receiver *MaintenanceService) GetBranchByCity(CityID int)(Branches []models.Branch, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetBranchByCity, CityID)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next(){
		Branch := models.Branch{}
		errRole := rows.Scan(&Branch.ID, &Branch.Address, &Branch.CityID)
		if errRole != nil {
			return
		}
		Branches = append(Branches, Branch)
	}
	return Branches, nil
}

func (receiver *MaintenanceService) AddPurpose(Purpose models.Purpose) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddPurpose, Purpose.Name)
	if err != nil {
		return
	}
	return nil
}

func (receiver *MaintenanceService) GetPurposes()(Purposes []models.Purpose, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetPurposes)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next(){
		Purpose := models.Purpose{}
		errRole := rows.Scan(&Purpose.ID, &Purpose.Name)
		if errRole != nil {
			return
		}
		Purposes = append(Purposes, Purpose)
	}
	return Purposes, nil
}

func (receiver *MaintenanceService) AddTime(Time models.Time) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddTime, Time.Name)
	if err != nil {
		return
	}
	return nil
}

func (receiver *MaintenanceService) GetTimes()(Times []models.Time, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetTimes)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next(){
		Time := models.Time{}
		errRole := rows.Scan(&Time.ID, &Time.Name)
		if errRole != nil {
			return
		}
		Times = append(Times, Time)
	}
	return Times, nil
}

func (receiver *MaintenanceService) AddTerminal(Terminal models.Terminal, Claims *tokens.Claims) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddTerminal, Terminal.TerminalNumber, Terminal.CityID, Terminal.BranchID, Claims.UserID)
	if err != nil {
		return
	}
	return nil
}