package app

import (
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"queue/api/services"
	"queue/databases/postgres"
)

type MainServer struct {
	pool *pgxpool.Pool
	router *httprouter.Router
	Initilize *postgres.DBPostgres
	userService *services.UserService
	maintenanceService *services.MaintenanceService
	queueService *services.QueueService
}

func NewMainServer(pool *pgxpool.Pool, router *httprouter.Router, initilize *postgres.DBPostgres, userService *services.UserService, maintenanceService *services.MaintenanceService, queueService *services.QueueService) *MainServer {
	return &MainServer{pool: pool, router: router, Initilize: initilize, userService: userService, maintenanceService: maintenanceService, queueService: queueService}
}


func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request){

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, RefreshToken")

	server.router.ServeHTTP(writer, request)
}

func(server *MainServer) Start() {
	err := server.Initilize.DbInit()
	if err != nil {
		panic("server don't created")
	}
	server.InitRoutes()
}