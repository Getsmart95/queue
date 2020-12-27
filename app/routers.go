package app

import (
	"queue/api/middlewares"
)

func (server *MainServer) InitRoutes() {
	// Authenticate
	server.router.POST("/api/register", server.RegisterHandler)
	server.router.GET("/api/login", server.LoginHandler)
	// Users
	server.router.GET("/api/roles", middlewares.JWT()(server.GetRolesHandler))
	server.router.POST("/api/users/addUser", middlewares.GuardRole(server.userService)(server.AddUserHandler))
	server.router.PUT("/api/users/updateUser", middlewares.JWT()(server.UpdateUserHandler))
	//server.router.GET("/api/users/getUserById/:{user_id}", server.GetUserByIdHandler)
	// Cities
	server.router.POST("/api/cities/addCity", middlewares.GuardRole(server.userService)(server.AddCity))
	server.router.GET("/api/cities", middlewares.JWT()(server.GetAllCities))
	// Branches
	server.router.POST("/api/branches/addBranch", middlewares.GuardRole(server.userService)(server.AddBranchHandler))
	server.router.GET("/api/branches/:city_id", middlewares.JWT()(server.GetBranchByCityHandler))
	// Times
	server.router.POST("/api/times/addTime", middlewares.GuardRole(server.userService)(server.AddTimesHandler))
	server.router.GET("/api/times", middlewares.JWT()(server.GetTimes))
	// Purposes
	server.router.POST("/api/purposes/addPurpose", middlewares.GuardRole(server.userService)(server.AddPurposeHandler))
	server.router.GET("/api/purposes", middlewares.JWT()(server.GetPurposes))
	// Queue
	server.router.POST("/api/queue/addQueue", middlewares.GuardRole(server.userService)(server.AddQueueHandler))
	server.router.GET("/api/queues/qetByDate/:date", middlewares.JWT()(server.GetQueuesByDateHandler))
	server.router.GET("/api/queues/qetByTime/:time_id", middlewares.JWT()(server.GetQueuesByTimeHandler))
	server.router.GET("/api/queues/getByStatus/:status", middlewares.JWT()(server.GetQueuesByStatusHandler))
	server.router.GET("/api/queues/getByUser/:user_id", middlewares.JWT()(server.GetQueuesByUserHandler))
	server.router.PUT("/api/queues/updateQueue", middlewares.JWT()(server.UpdateQueueHandler))
	server.router.PUT("/api/queues/changeStatus/:queue_id", middlewares.JWT()(server.QueueChangeStatusHandler))
	// Notification
	//server.router.POST("/api/sms/send", server.SendNotificationHandler)
}
