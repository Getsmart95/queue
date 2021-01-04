package app

import (
	"encoding/json"
	"fmt"
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
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("Invalid json")
		return
	}

	result, err := server.userService.CheckUser(requestBody.Login)
	if err != nil || result == false {
		responseBody.Ok = false
		responseBody.Message = "Пользователь не найден"
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(responseBody)
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
		json.NewEncoder(writer).Encode(responseBody)
		return
	}
	return
}

func (server *MainServer) RegisterHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseToken
	var responseStatus models.ResponseStatus
	var Status models.CredentialStatus

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("Invalid json")
		return
	}

	if len(requestBody.Password) < 8 {
		responseStatus.Ok = false
		responseStatus.Message = "Password is too short. Minimum 8 symbols"
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(responseStatus)
		return
	}
	err = server.userService.Registration(requestBody)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Status)
		return
	}
	//	Add role "user"
	roleID := 3
	err = server.userService.AddUserRole(roleID, requestBody.Login)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, responseUser, err := server.userService.GetUserByLogin(requestBody.Login)

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

func (server *MainServer) GetRolesHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
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

func (server *MainServer) AddManagerHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseStatus

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestBody.Password) < 8 {
		responseBody.Ok = false
		responseBody.Message = "Password is too short. Minimum 8 symbols"
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(responseBody)
		return
	}

	err = server.userService.AddManager(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Add role "manager"
	//roleID, _ := strconv.Atoi(params.ByName("role_id")) // Можно использовать если необходимо будет добавлять Главного администратора
	roleID := 2
	err = server.userService.AddUserRole(roleID, requestBody.Login)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	_, _, err = server.userService.GetUserByLogin(requestBody.Login)

	responseBody.Ok = true
	responseBody.Message = "Пользователь успешно добавлен"
	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) UpdateManagerHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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

func (server *MainServer) GetAllCities(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

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

func (server *MainServer) GetBranchByCityHandler(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {
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

func (server *MainServer) GetPurposes(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

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

func (server *MainServer) GetTimes(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

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

func (server *MainServer) AddQueueOnlineHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.RequestQueue
	var responseBody models.ResponseStatus
	claims := middlewares.JWT(writer, request, params)
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	queueCode, err := server.queueService.GetLastQueueByDate(requestBody.Date)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.queueService.AddQueueOnline(requestBody, queueCode, claims)
	fmt.Println(err)
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

func (server *MainServer) GetQueuesByDateHandler(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {

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

func (server *MainServer) GetQueuesByTimeHandler(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {

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
	var RequestBody models.RequestDate
	Status := params.ByName("status")
	err := json.NewDecoder(request.Body).Decode(&RequestBody)
	if err != nil {
		json.NewEncoder(writer).Encode("Invalid json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(Status)
	fmt.Println(RequestBody)
	queues, err := server.queueService.GetQueuesByStatus(Status, RequestBody.Date)
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

func (server *MainServer) GetQueuesByUserHandler(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {

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

func (server *MainServer) UpdateQueueHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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

func (server *MainServer) NotificationHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
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

