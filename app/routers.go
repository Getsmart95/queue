package app

import (
	"reflect"
)

func (server *MainServer) InitRoutes() {
	//handler := jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), server.secret)(server.RemoveUserByLoginHandler)

	server.router.GET("/api/roles", server.GetRolesHandler)
	server.router.POST("/api/users/addUser", server.AddUserHandler)
	server.router.GET("/api/users/getUserById/:{user_id}", server.GetUserByIdHandler)
	server.router.POST("/api/cities/addCity", server.AddCity)
	server.router.GET("/api/cities/getCities", server.GetAllCities)
	server.router.POST("/api/branches/addBranch", server.AddBranchHandler)
	server.router.GET("/api/branches/getBranch/:{city_id}", server.GetBranchByCityHandler)
	server.router.POST("/api/times/addTime", server.AddTimesHandler)
	server.router.GET("/api/times/gettimes", server.GetTimes)
	//	server.router.DELETE(removeByLogin, server.RemoveUserByLoginHandler)
	//	server.router.GET(getUserList, server.GetUserListHandler)
	//	server.router.POST(saveNewUser, server.SaveNewUserHandler)
}
