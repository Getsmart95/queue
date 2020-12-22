package postgres

import (
	"context"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

type DBPostgres struct {
	pool *pgxpool.Pool
}

func NewDBInit(pool *pgxpool.Pool) *DBPostgres {
	return &DBPostgres{pool: pool}
}

func (receiver *DBPostgres) DbInit() (err error) {
	var DDLs = []string {
		CreateUsersTable,
		CreateBranchesTable,
		CreateCitiesTable,
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
			log.Fatal("err: #{err}\n")
			return err
		}
	}
	return nil

}
