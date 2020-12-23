package app

import (
	//"reflect"
)

func (server *MainServer) InitRoutes() {
	//handler := jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), server.secret)(server.RemoveUserByLoginHandler)
	// Authenticate
	server.router.POST("/api/register", server.RegisterHandler)
	server.router.GET("/api/login", server.LoginHandler)
	// Users
	server.router.GET("/api/roles", server.GetRolesHandler)
	server.router.POST("/api/users/addUser", server.AddUserHandler)
	//server.router.GET("/api/users/getUserById/:{user_id}", server.GetUserByIdHandler)
	// Cities
	server.router.POST("/api/cities/addCity", server.AddCity)
	server.router.GET("/api/cities/getCities", server.GetAllCities)
	// Branches
	server.router.POST("/api/branches/addBranch", server.AddBranchHandler)
	server.router.GET("/api/branches/getBranch/:{city_id}", server.GetBranchByCityHandler)
	// Times
	server.router.POST("/api/times/addTime", server.AddTimesHandler)
	server.router.GET("/api/times/gettimes", server.GetTimes)
}
