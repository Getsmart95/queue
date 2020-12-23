package app

import (
	"encoding/json"
	"fmt"
	"queue/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)
const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) RegisterHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.User
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.userService.Registration(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	userID, err := server.userService.GetUserByLogin(requestBody.Login)
	//	Add role "user"
	roleID := 3
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

func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.User
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	fmt.Println(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	//json.NewEncoder(writer).Encode(err)
	Login, Password, err := server.userService.Authentication(requestBody)
	fmt.Println(Login, Password,err)
	if err != nil && Login == false {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Login)
	}

	if Password == false {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Password)
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

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
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetAllCities(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cities, err := server.maintenanceService.GetAllCities()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(cities)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
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
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetBranchByCityHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cityID, _ := strconv.Atoi(params.ByName("city_id"))
	branches, err := server.maintenanceService.GetBranchByCity(cityID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(branches)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddPurposeHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.Purpose
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddPurpose(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetPurposes(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	purposes, err := server.maintenanceService.GetPurposes()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(purposes)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}


func (server *MainServer) AddTimesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.Time
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddTime(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetTimes(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	times, err := server.maintenanceService.GetTimes()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(times)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}