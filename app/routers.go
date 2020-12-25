package app

import "queue/middleware"

func (server *MainServer) InitRoutes() {
	// Authenticate
	server.router.POST("/api/register", server.RegisterHandler)
	server.router.GET("/api/login", server.LoginHandler)
	// Users
	server.router.GET("/api/roles", middleware.JWT()(server.GetRolesHandler))
	server.router.POST("/api/users/addUser", server.AddUserHandler)
	//server.router.GET("/api/users/getUserById/:{user_id}", server.GetUserByIdHandler)
	// Cities
	server.router.POST("/api/cities/addCity", server.AddCity)
	server.router.GET("/api/cities", server.GetAllCities)
	// Branches
	server.router.POST("/api/branches/addBranch", server.AddBranchHandler)
	server.router.GET("/api/branches/:city_id", server.GetBranchByCityHandler)
	// Times
	server.router.POST("/api/times/addTime", server.AddTimesHandler)
	server.router.GET("/api/times", server.GetTimes)
	// Purposes
	server.router.POST("/api/purposes/addPurpose", server.AddPurposeHandler)
	server.router.GET("/api/purposes", server.GetPurposes)
	// Queue
	server.router.POST("/api/queue/addQueue", server.AddQueueHandler)
	server.router.GET("/api/queues/:date", server.GetQueuesHandler)
}
