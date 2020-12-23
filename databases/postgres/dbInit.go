package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
)

type DBPostgres struct {
	pool *pgxpool.Pool
}

func NewDBInit(pool *pgxpool.Pool) *DBPostgres {
	return &DBPostgres{pool: pool}
}

func (receiver *DBPostgres) DbInit() (err error) {
	var DDLs = []string{
		CreateUsersTable,
		CreateCitiesTable,
		CreateBranchesTable,
		CreatePurposesTable,
		CreateTimesTable,
		CreateRolesTable,
		CreateUserRoleTable,
		CreateQueuesTable,
		RolesSeeder,
		AddAdmin}

	for _, ddl :=range DDLs {
		_, err := receiver.pool.Exec(context.Background(), ddl)
		if err != nil {
			fmt.Println("errs: #{err}\n")
			return err
		}
	}
	return nil

}
