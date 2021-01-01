package app

import (
	"queue/api/middlewares"
)

func (server *MainServer) InitRoutes() {

	// Authenticate
	server.router.POST("/api/checkUser", server.CheckUserHandler) // Humo online
	server.router.POST("/api/register", server.RegisterHandler) // Humo online
	server.router.POST("/api/login", server.LoginHandler) // For ADMIN & Humo online
	// Users
	server.router.GET("/api/roles", middlewares.JWTAuth()(server.GetRolesHandler))
	server.router.POST("/api/users/addManager", middlewares.GuardRole(server.userService)(server.AddManagerHandler))
	server.router.PUT("/api/users/updateManager", middlewares.JWTAuth()(server.UpdateManagerHandler))
	//server.router.GET("/api/users/getUserById/:{user_id}", server.GetUserByIdHandler)
	// Cities
	server.router.POST("/api/cities/addCity", middlewares.GuardRole(server.userService)(server.AddCity))
	server.router.GET("/api/cities", middlewares.JWTAuth()(server.GetAllCities))
	// Branches
	server.router.POST("/api/branches/addBranch", middlewares.GuardRole(server.userService)(server.AddBranchHandler))
	server.router.GET("/api/branches/:city_id", middlewares.JWTAuth()(server.GetBranchByCityHandler))
	// Times
	server.router.POST("/api/times/addTime", middlewares.GuardRole(server.userService)(server.AddTimesHandler))
	server.router.GET("/api/times", middlewares.JWTAuth()(server.GetTimes))
	// Purposes
	server.router.POST("/api/purposes/addPurpose", middlewares.GuardRole(server.userService)(server.AddPurposeHandler))
	server.router.GET("/api/purposes", middlewares.JWTAuth()(server.GetPurposes))
	// Queue
	server.router.POST("/api/queue/addQueue", middlewares.GuardRole(server.userService)(server.AddQueueHandler)) // Terminal
	server.router.GET("/api/queues/getByDate/:date", middlewares.JWTAuth()(server.GetQueuesByDateHandler))
	server.router.GET("/api/queues/getByTime/:time_id", middlewares.JWTAuth()(server.GetQueuesByTimeHandler))
	server.router.GET("/api/queues/getByStatus/:status", middlewares.JWTAuth()(server.GetQueuesByStatusHandler))
	server.router.GET("/api/queues/getByUser/:user_id", middlewares.JWTAuth()(server.GetQueuesByUserHandler))
	server.router.PUT("/api/queues/updateQueue", middlewares.JWTAuth()(server.UpdateQueueHandler))
	server.router.PUT("/api/queues/changeStatus/:queue_id", middlewares.JWTAuth()(server.QueueChangeStatusHandler))
	// Notification
	server.router.POST("/api/sms/send", server.NotificationHandler)
	// Terminal
	server.router.POST("/api/terminal/addTerminal", middlewares.CORS(middlewares.GuardRole(server.userService)(server.AddTerminalHandler)))
}
