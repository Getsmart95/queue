package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
	"queue/databases/postgres"
	"queue/models"
	"queue/tokens"
	"strconv"
)

type QueueService struct {
	pool *pgxpool.Pool
}

func NewQueueService(pool *pgxpool.Pool) *QueueService {
	return &QueueService{pool: pool}
}

func (receiver *QueueService) AddQueue(Queue models.RequestTerminal, LastQueueCode int, Claims *tokens.Claims)(err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	var Terminal models.Terminal
	err = conn.QueryRow(context.Background(), postgres.GetTerminalByUserID, Claims.UserID).Scan(
		&Terminal.ID,
		&Terminal.TerminalNumber,
		&Terminal.CityID,
		&Terminal.BranchID,
		&Terminal.UserID)
	defer conn.Release()
	Status := "Pending"

	QueueCode := LastQueueCode + 1
	fmt.Printf("%010d", QueueCode)
	_, err = conn.Exec(context.Background(), postgres.AddQueue,
		strconv.Itoa(QueueCode),
		Terminal.UserID,
		Terminal.TerminalNumber,
		Terminal.CityID,
		Terminal.BranchID,
		Queue.PurposeID,
		Status,
		Queue.Date)

	if err != nil {
		return
	}
	return nil
}

func (receiver *QueueService) GetQueuesByDate(Date string)(queues []models.Queue, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetQueuesByDate, Date)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		Queue := models.Queue{}
		errQueue := rows.Scan(
			&Queue.ID,
			&Queue.QueueCode,
			&Queue.TerminalID,
			&Queue.UserID,
			&Queue.CityID,
			&Queue.BranchID,
			&Queue.PurposeID,
			&Queue.TimeID,
			&Queue.Status,
			&Queue.Date,
			&Queue.StartAt,
			&Queue.FinishAt,
			&Queue.CreatedAt)

		if errQueue != nil {
			return
		}
		queues = append(queues, Queue)
	}
	return queues, nil
}

func (receiver *QueueService) GetLastQueueByDate(Date string)(queue int, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(), postgres.GetLastQueueByDate, Date).Scan(
		&queue)
	fmt.Println(err)

	if err != nil {
		return
	}
fmt.Println(queue)

	return queue, nil
}


func (receiver *QueueService) GetQueuesByTime(TimeID int)(queues []models.Queue, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetQueuesByTime, TimeID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		Queue := models.Queue{}
		errQueue := rows.Scan(
			&Queue.ID,
			&Queue.QueueCode,
			&Queue.UserID,
			&Queue.CityID,
			&Queue.BranchID,
			&Queue.PurposeID,
			&Queue.TimeID,
			&Queue.Status,
			&Queue.Date,
			&Queue.StartAt,
			&Queue.FinishAt,
			&Queue.CreatedAt)
		if errQueue != nil {
			return
		}
		queues = append(queues, Queue)
	}
	return queues, nil
}

func (receiver *QueueService) GetQueuesByStatus(Status string)(queues []models.Queue, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetQueuesByStatus, Status)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		Queue := models.Queue{}
		errQueue := rows.Scan(
			&Queue.ID,
			&Queue.QueueCode,
			&Queue.UserID,
			&Queue.TerminalID,
			&Queue.CityID,
			&Queue.BranchID,
			&Queue.PurposeID,
			&Queue.TimeID,
			&Queue.Status,
			&Queue.Date,
			&Queue.StartAt,
			&Queue.FinishAt,
			&Queue.CreatedAt)
		if errQueue != nil {
			return
		}
		queues = append(queues, Queue)
	}
	return queues, nil
}

func (receiver *QueueService) GetQueuesByUser(UserID int)(queues []models.Queue, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), postgres.GetQueuesByUser, UserID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		Queue := models.Queue{}
		errQueue := rows.Scan(
			&Queue.ID,
			&Queue.QueueCode,
			&Queue.UserID,
			&Queue.TerminalID,
			&Queue.CityID,
			&Queue.BranchID,
			&Queue.PurposeID,
			&Queue.TimeID,
			&Queue.Status,
			&Queue.Date,
			&Queue.StartAt,
			&Queue.FinishAt,
			&Queue.CreatedAt)
		if errQueue != nil {
			return
		}
		queues = append(queues, Queue)
	}
	return queues, nil
}

func (receiver *QueueService) UpdateQueue(Queue models.Queue)(err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.UpdateQueue,
		Queue.QueueCode,
		Queue.UserID,
		Queue.TerminalID,
		Queue.CityID,
		Queue.BranchID,
		Queue.PurposeID,
		Queue.TimeID,
		Queue.Status,
		Queue.Date,
		Queue.ID)

	if err != nil {
		return
	}
	return nil
}

func (receiver *QueueService) QueueChangeStatus(QueueID int, Status models.RequestStatus)(err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.QueueChangeStatus,Status.Status, QueueID)

	if err != nil {
		return
	}
	return nil
}

