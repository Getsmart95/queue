package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"queue/api/services"
	"queue/app"
	"queue/databases/postgres"
)

var (
	host = flag.String("host", "127.0.0.1", "Server host")
	port = flag.String("port", "8081", "Server port")
	dsn  = flag.String("dsn", "postgresql://root@localhost:5432/postgres?sslmode=disable", "Postgres DSN")
)

func main() {
	flag.Parse()
	address := net.JoinHostPort(*host, *port)
	router := httprouter.New()
	pool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		log.Printf("%e", err)
	}

	userService := services.NewUserService(pool)
	maintenanceService := services.NewMaintenanceService(pool)
	queueService := services.NewQueueService(pool)
	dbInit := postgres.NewDBInit(pool)

	server := app.NewMainServer(pool, router,  dbInit, userService, maintenanceService, queueService)

	server.Start()
	fmt.Println(address)
	panic(http.ListenAndServe(address, server))


}
