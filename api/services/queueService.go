package services

import (
	"context"
	"fmt"
	"queue/databases/postgres"
	"queue/models"
	"github.com/jackc/pgx/pgxpool"
)

type QueueService struct {
	pool *pgxpool.Pool
}

func NewQueueService(pool *pgxpool.Pool) *QueueService {
	return &QueueService{pool: pool}
}

func (receiver *QueueService) AddQueue(Queue models.Queue)(err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		return
	}

	defer conn.Release()

	_, err = conn.Exec(context.Background(), postgres.AddQueue,
		Queue.QueueCode,
		Queue.UserID,
		Queue.CityID,
		Queue.BranchID,
		Queue.PurposeID,
		Queue.TimeID,
		Queue.Status,
		Queue.Date)

	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

func (receiver *QueueService) GetQueues(Date string)(queues []models.Queue, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		return
	}
	defer conn.Release()
	fmt.Println(Date)
	rows, err := conn.Query(context.Background(), postgres.GetQueuesByDate, Date)
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
			&Queue.Date)
		if errQueue != nil {
			return
		}
		queues = append(queues, Queue)
	}
	fmt.Println(rows)
	fmt.Println(queues)
	return queues, nil
}
