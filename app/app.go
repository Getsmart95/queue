package app

import (
	"net/http"
	"queue/api/services"
	"queue/databases/postgres"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type MainServer struct {
	pool *pgxpool.Pool
	router *httprouter.Router
	Initilize *postgres.DBPostgres
	userService *services.UserService
	maintenanceService *services.MaintenanceService
}

func NewMainServer(pool *pgxpool.Pool, router *httprouter.Router, initilize *postgres.DBPostgres, userService *services.UserService, maintenanceService *services.MaintenanceService) *MainServer {
	return &MainServer{pool: pool, router: router, Initilize: initilize, userService: userService, maintenanceService: maintenanceService}
}

func (server *MainServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	server.router.ServeHTTP(w, r)
}

func(server *MainServer) Start() {
	err := server.Initilize.DbInit()
	if err != nil {
		panic("server don't created")
	}
	server.InitRoutes()
}