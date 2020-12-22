package app

import (
	"encoding/json"
	"queue/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)
const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) GetRolesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	roles, err := server.userService.GetAllRoles()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(roles)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return

}

func (server *MainServer) AddUserHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.User
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.userService.AddUser(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	userID, err := server.userService.GetUserByLogin(requestBody.Login)
	roleID, _ := strconv.Atoi(params.ByName("role_id"))
	err = server.userService.AddUserRole(userID, roleID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddCity(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.City
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddCity(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (server *MainServer) GetAllCities(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
}
func (server *MainServer) AddBranchHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.Branch
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddBranch(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (server *MainServer) GetBranchByCity(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
