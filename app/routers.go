package app

import (
	"reflect"
)

func (server *MainServer) InitRoutes() {
	//handler := jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), server.secret)(server.RemoveUserByLoginHandler)

	server.router.GET("/api/getAllRoles", server.GetRolesHandler)
	server.router.POST("/api/addUser", server.AddUserHandler)
	//	server.router.DELETE(removeByLogin, server.RemoveUserByLoginHandler)
	//	server.router.GET(getUserList, server.GetUserListHandler)
	//	server.router.POST(saveNewUser, server.SaveNewUserHandler)
}
