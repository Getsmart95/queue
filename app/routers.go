package app

import "queue/middleware"

func (server *MainServer) InitRoutes() {
	// Authenticate
	server.router.POST("/api/register", server.RegisterHandler)
	server.router.GET("/api/login", server.LoginHandler)
	// Users
	server.router.GET("/api/roles", middleware.JWT()(server.GetRolesHandler))
	server.router.POST("/api/users/addUser", middleware.JWT()(server.AddUserHandler))
	server.router.PUT("/api/users/updateUser", server.UpdateUserHandler)
	//server.router.GET("/api/users/getUserById/:{user_id}", server.GetUserByIdHandler)
	// Cities
	server.router.POST("/api/cities/addCity", middleware.JWT()(server.AddCity))
	server.router.GET("/api/cities", server.GetAllCities)
	// Branches
	server.router.POST("/api/branches/addBranch", middleware.JWT()(server.AddBranchHandler))
	server.router.GET("/api/branches/:city_id", server.GetBranchByCityHandler)
	// Times
	server.router.POST("/api/times/addTime", middleware.JWT()(server.AddTimesHandler))
	server.router.GET("/api/times", server.GetTimes)
	// Purposes
	server.router.POST("/api/purposes/addPurpose", middleware.JWT()(server.AddPurposeHandler))
	server.router.GET("/api/purposes", server.GetPurposes)
	// Queue
	server.router.POST("/api/queue/addQueue", middleware.JWT()(server.AddQueueHandler))
	server.router.GET("/api/queues/qetByDate/:date", server.GetQueuesByDateHandler)
	server.router.GET("/api/queues/qetByTime/:time_id", server.GetQueuesByTimeHandler)
	server.router.GET("/api/queues/getByStatus/:status", server.GetQueuesByStatusHandler)
	server.router.GET("/api/queues/getByUser/:user_id", server.GetQueuesByUserHandler)
	server.router.PUT("/api/queues/updateQueue", middleware.JWT()(server.UpdateQueueHandler))
	server.router.PUT("/api/queues/changeStatus/:queue_id", middleware.JWT()(server.QueueChangeStatusHandler))
	// Notification
	//server.router.POST("/api/sms/send", server.SendNotificationHandler)
}
