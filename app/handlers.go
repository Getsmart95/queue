package app

import (
	"queue/tokens"
	"encoding/json"
	"fmt"
	"queue/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)
const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) RegisterHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseToken
	var Status models.CredentialStatus
	writer.Header().Set(contentType, value)
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.userService.Registration(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Status)
		return
	}
	user, err := server.userService.GetUserByLogin(requestBody.Login)
	//	Add role "user"
	roleID := 3
	err = server.userService.AddUserRole(user.ID, roleID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	token, expiredIn := tokens.SetToken(requestBody.Login, requestBody.Password)

	responseBody.Ok = true
	responseBody.Token = token
	responseBody.ExpiredIn = expiredIn
	responseBody.Status = 200
	responseBody.User = user

	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

}

func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var User models.User
	var Status models.CredentialStatus
	var requestBody models.User
	var responseBody models.ResponseToken
	writer.Header().Set(contentType, value)

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	fmt.Println(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	Status.Login ,Status.Password, User, err = server.userService.Authentication(requestBody)
	fmt.Println(Status.Login, Status.Password,err)
	if err != nil && Status.Login == false {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Status)
		return
	}

	if Status.Password == false {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Status)
		return
	}

	token, expiredIn := tokens.SetToken(requestBody.Login, requestBody.Password)

	responseBody.Ok = true
	responseBody.Token = token
	responseBody.ExpiredIn = expiredIn
	responseBody.Status = 200
	responseBody.User = User

	err = json.NewEncoder(writer).Encode(responseBody)
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
	user, err := server.userService.GetUserByLogin(requestBody.Login)
	roleID, _ := strconv.Atoi(params.ByName("role_id"))
	err = server.userService.AddUserRole(user.ID, roleID)
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
	writer.Header().Set(contentType, value)
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
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetBranchByCityHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(contentType, value)
	cityID, _ := strconv.Atoi(params.ByName("city_id"))
	fmt.Println(params.ByName("city_id"))
	branches, err := server.maintenanceService.GetBranchByCity(cityID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

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
	fmt.Println(requestBody)
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

func (server *MainServer) AddQueueHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.Queue
	//var responseBody models.ResponseToken
	writer.Header().Set(contentType, value)
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.queueService.AddQueue(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (server *MainServer) GetQueuesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(contentType, value)
	Date := params.ByName("date")
	queues, err := server.queueService.GetQueues(Date)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(queues)
	err = json.NewEncoder(writer).Encode(queues)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}