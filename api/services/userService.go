package services

import (
	"context"
	"errors"
	"github.com/jackc/pgx/pgxpool"
	"log"
	"queue/databases/postgres"
	"queue/models"
)

type UserService struct {
	pool *pgxpool.Pool
}

func NewUserService(pool *pgxpool.Pool) *UserService{
	if pool == nil{
		log.Println(errors.New("test"))
	}
	return &UserService{pool: pool}
}

func (receiver *UserService) GetAllRoles() (Roles []models.Role, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("can't get connection")
		return
	}

	defer conn.Release()
	rows, err := conn.Query(context.Background(), postgres.GetAllRoles)
	if err != nil {
		log.Fatal("no have roles")
		return
	}

	defer rows.Close()
	for rows.Next(){
		Role := models.Role{}
		errRole := rows.Scan(&Role.ID, &Role.Name, &Role.DisplayName, &Role.Description)
		if errRole != nil {
			log.Fatal("canr read #{errRole}")
			return
		}
		Roles = append(Roles, Role)
	}
	return Roles, nil
}

func (receiver *UserService) AddUser(User models.User) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("can't get connection")
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddUser, User.Name, User.Surname, User.Login, User.Password, User.Email, User.Phone, User.Status)
	if err != nil {
		log.Fatal("Cant add user")
	}
	return nil
}

func (receiver *UserService) AddUserRole(userID int, roleID int) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("can't get connection")
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddUserRole, roleID, userID)
}

func (receiver *UserService) GetUserByLogin(userLogin string) (userID int, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("can't get connection")
	}
	defer conn.Release()
	var UserID int
	err = conn.QueryRow(context.Background(), postgres.GetUserByLogin, userLogin).Scan(&UserID)
	if err != nil {
		log.Fatal("Cant add user")
	}
	return UserID, nil
}
