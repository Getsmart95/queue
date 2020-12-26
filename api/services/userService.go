package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"queue/databases/postgres"
	"queue/models"
)

type UserService struct {
	pool *pgxpool.Pool
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{pool: pool}
}

func (receiver *UserService) Registration(User models.User) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()
	password := MakeHash(User.Password)
	_, err = conn.Exec(context.Background(), postgres.AddUser, User.Name, User.Surname, User.Login, password, User.Email, User.Phone, User.Status)
	if err != nil {
		return
	}

	return nil
}

func (receiver *UserService) Authentication(User models.User) (Login bool, Password bool, user models.User, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(), postgres.GetUserByLogin, User.Login).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.Status,
		&user.CreatedAt)
	fmt.Println(MakeHash(User.Password))
	fmt.Println(user.Password)
	fmt.Println(CompareHashWithPassword(user.Password,User.Password))
	if err != nil {
		return false, false, user, err
	}
	errHash := CompareHashWithPassword(user.Password, User.Password)
	if errHash != nil {
		return true, false, user, err
	}

	return true, true, user, nil
}

func CompareHashWithPassword(HashedPassword string, Password string) error {
	result := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(Password))
	return result
}

func MakeHash(Password string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Can't hash password")
	}
	return string(password)
}

func (receiver *UserService) GetAllRoles() (Roles []models.Role, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()
	rows, err := conn.Query(context.Background(), postgres.GetAllRoles)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next(){
		Role := models.Role{}
		errRole := rows.Scan(&Role.ID, &Role.Name, &Role.DisplayName, &Role.Description)
		if errRole != nil {
			return
		}
		Roles = append(Roles, Role)
	}
	return Roles, nil
}

func (receiver *UserService) AddUser(User models.User) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()
	password := MakeHash(User.Password)
	_, err = conn.Exec(context.Background(), postgres.AddUser, User.Name, User.Surname, User.Login, password, User.Email, User.Phone, User.Status)
	if err != nil {
		return
	}
	return nil
}

func (receiver *UserService) AddUserRole(userID int, roleID int) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddUserRole, roleID, userID)
	return nil
}

func (receiver *UserService) GetUserByLogin(userLogin string) (user models.User, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	err = conn.QueryRow(context.Background(), postgres.GetUserByLogin, userLogin).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.Status,
		&user.CreatedAt)
	if err != nil {
		return
	}
	return user, nil
}

func (receiver *UserService) UpdateUser(User models.User)(err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.UpdateUser,
		User.Name,
		User.Surname,
		User.Email,
		User.Phone,
		User.Status,
		User.ID)

	if err != nil {
		return
	}
	return nil
}
