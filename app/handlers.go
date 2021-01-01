package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"queue/api/middlewares"
	"queue/models"
	"queue/tokens"
	"strconv"
)

func (server *MainServer) CheckUserHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.RequestLogin
	var responseBody models.ResponseStatus
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := server.userService.CheckUser(requestBody.Login)
	if err != nil {
		json.NewEncoder(writer).Encode("Undefined error #{err}")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if result == false {
		responseBody.Ok = false
		responseBody.Message = "Пользователь не существует"
		json.NewEncoder(writer).Encode(responseBody)
		return
	} else if result == true {
		responseBody.Ok = true
		responseBody.Message = "Пользователь существует"
	}
}

func (server *MainServer) RegisterHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseToken
	var Status models.CredentialStatus

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
	userID, responseUser, err := server.userService.GetUserByLogin(requestBody.Login)
	//	Add role "user"
	roleID := 3
	err = server.userService.AddUserRole(userID, roleID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	token := tokens.SetToken(userID, responseUser.Login)

	responseBody.Ok = true
	responseBody.Token = token
	responseBody.User = responseUser

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


	err := json.NewDecoder(request.Body).Decode(&requestBody)

	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var userID int
	Status.Login ,Status.Password, userID, User, err = server.userService.Authentication(requestBody)

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

	token := tokens.SetToken(userID, User.Login)

	responseBody.Ok = true
	responseBody.Token = token
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
	err = json.NewEncoder(writer).Encode(roles)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddManagerHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseStatus

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.userService.AddUser(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	userID, _, err := server.userService.GetUserByLogin(requestBody.Login)
	//roleID, _ := strconv.Atoi(params.ByName("role_id")) // Можно использовать если необходимо будет добавлять Главного администратора
	// Add role "manager"
	roleID := 2
	err = server.userService.AddUserRole(userID, roleID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	responseBody.Ok = true
	responseBody.Message = "Пользователь успешно добавлен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) UpdateManagerHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseStatus

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.userService.UpdateUser(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	responseBody.Ok = true
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

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddCity(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
	responseBody.Message = "Город успешно добавлен"
	err = json.NewEncoder(writer).Encode(responseBody)
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

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddBranch(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
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

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddPurpose(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
	responseBody.Message = "Цель визита успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
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

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.maintenanceService.AddTime(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
	responseBody.Message = "Время успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
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

	err = json.NewEncoder(writer).Encode(times)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddQueueHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.RequestTerminal
	var responseBody models.ResponseStatus
	claims := middlewares.JWT(writer, request, params)
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	queueCode, err := server.queueService.GetLastQueueByDate(requestBody.Date)
	err = server.queueService.AddQueue(requestBody, queueCode, claims)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	responseBody.Ok = true
	responseBody.Message = "Запись в очередь успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetQueuesByDateHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

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

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.queueService.UpdateQueue(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
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

	QueueID, _ := strconv.Atoi(params.ByName("queue_id"))
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.queueService.QueueChangeStatus(QueueID, requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody.Ok = true
	responseBody.Message = "Запись в очередь успешно добавлена"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) NotificationHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var responseBody models.ResponseStatus

	responseBody.Ok = true
	responseBody.Message = "Уведомление успешно отправлено"
	err := json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) AddTerminalHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.Terminal
	var responseBody models.ResponseStatus

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	claims := middlewares.JWT(writer, request, params)
	err = server.maintenanceService.AddTerminal(requestBody, claims)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	responseBody.Ok = true
	responseBody.Message = "Новый терминал успешно добавлен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

