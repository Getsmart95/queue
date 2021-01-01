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

func (receiver *UserService) CheckUser(Login string) (result bool, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	var exists string
	err = conn.QueryRow(context.Background(), postgres.CheckUser, Login).Scan(&exists)

	if err != nil {
		return false, err
	}
	if exists == "" {
		return false, nil
	} else {
		return true, nil
	}

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

func (receiver *UserService) Authentication(User models.User) (login bool, password bool, userID int, responseUser models.ResponseUser, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	var user models.UserWithRole
	err = conn.QueryRow(context.Background(), postgres.GetUserByLogin, User.Login).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.Role,
		&user.Status,
		&user.CreatedAt)
	if err != nil {
		return false, false, userID, responseUser, err
	}
	fmt.Println(user)
	responseUser = GetResponseUser(user)

	errHash := CompareHashWithPassword(user.Password, User.Password)
	if errHash != nil {
		return true, false, userID, responseUser, err
	}

	return true, true, user.ID, responseUser, nil
}

func GetResponseUser(User models.UserWithRole) (responseUser models.ResponseUser) {

	responseUser.Name = User.Name
	responseUser.Surname = User.Surname
	responseUser.Login = User.Login
	responseUser.Email = User.Email
	responseUser.Phone = User.Phone
	responseUser.Role = User.Role
	responseUser.Status = User.Status
	responseUser.CreatedAt = User.CreatedAt

	return responseUser
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

func (receiver *UserService) GetRoleByUser(Login string) (role models.JWTUserRole, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()
	err = conn.QueryRow(context.Background(), postgres.GetRoleByUser, Login).Scan(&role.RoleID, &role.UserID, &role.Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	return role, nil
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

func (receiver *UserService) GetUserByLogin(userLogin string) (userID int, responseUser models.ResponseUser, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	var user models.UserWithRole
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

	responseUser = GetResponseUser(user)
	return user.ID, responseUser, nil
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
