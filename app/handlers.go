package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"queue/models"
	"queue/tokens"
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
	var User models.ResponseUser
	var Status models.CredentialStatus
	var requestBody models.User
	var responseBody models.ResponseToken
	writer.Header().Set(contentType, value)

	err := json.NewDecoder(request.Body).Decode(&requestBody)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	Status.Login ,Status.Password, User, err = server.userService.Authentication(requestBody)

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
	writer.Header().Set(contentType, value)
	roles, err := server.userService.GetAllRoles()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	err = json.NewEncoder(writer).Encode(roles)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddUserHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseStatus
	writer.Header().Set(contentType, value)

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
	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Пользователь успешно добавлен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) UpdateUserHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseStatus
	writer.Header().Set(contentType, value)

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.userService.UpdateUser(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Профиль успешно обновлен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	return
}

func (server *MainServer) AddCity(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.City
	var responseBody models.ResponseStatus
	writer.Header().Set(contentType, value)

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

	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Город успешно добавлен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	return
}

func (server *MainServer) GetAllCities(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(contentType, value)

	cities, err := server.maintenanceService.GetAllCities()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(writer).Encode(cities)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddBranchHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.Branch
	var responseBody models.ResponseStatus
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

	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Отделение успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetBranchByCityHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cityID, _ := strconv.Atoi(params.ByName("city_id"))
	writer.Header().Set(contentType, value)

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
	var responseBody models.ResponseStatus
	writer.Header().Set(contentType, value)

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

	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Цель визита успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetPurposes(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set(contentType, value)

	purposes, err := server.maintenanceService.GetPurposes()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(writer).Encode(purposes)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}


func (server *MainServer) AddTimesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.Time
	var responseBody models.ResponseStatus
	writer.Header().Set(contentType, value)

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

	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Время успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetTimes(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set(contentType, value)

	times, err := server.maintenanceService.GetTimes()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(writer).Encode(times)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddQueueHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.Queue
	var responseBody models.ResponseStatus
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

	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Запись в очередь успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetQueuesByDateHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(contentType, value)

	Date := params.ByName("date")
	queues, err := server.queueService.GetQueuesByDate(Date)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(writer).Encode(queues)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetQueuesByTimeHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(contentType, value)

	TimeID, _ := strconv.Atoi(params.ByName("time_id"))
	queues, err := server.queueService.GetQueuesByTime(TimeID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(writer).Encode(queues)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetQueuesByStatusHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(contentType, value)

	Status := params.ByName("status")
	queues, err := server.queueService.GetQueuesByStatus(Status)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(writer).Encode(queues)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetQueuesByUserHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(contentType, value)

	UserID, _ := strconv.Atoi(params.ByName("user_id"))
	queues, err := server.queueService.GetQueuesByUser(UserID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(writer).Encode(queues)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) UpdateQueueHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.Queue
	var responseBody models.ResponseStatus
	writer.Header().Set(contentType, value)

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.queueService.UpdateQueue(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Запись успешно обновлен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) QueueChangeStatusHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.RequestStatus
	var responseBody models.ResponseStatus
	writer.Header().Set(contentType, value)

	QueueID, _ := strconv.Atoi(params.ByName("queue_id"))
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.queueService.QueueChangeStatus(QueueID, requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
	responseBody.Status = 200
	responseBody.Message = "Статус успешно изменен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

