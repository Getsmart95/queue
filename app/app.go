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
	userService *services.UserService
	Initilize *postgres.DBPostgres
}

func NewMainServer(pool *pgxpool.Pool, router *httprouter.Router, userService *services.UserService, initilize *postgres.DBPostgres) *MainServer {
	return &MainServer{pool: pool, router: router, userService: userService, Initilize: initilize}
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