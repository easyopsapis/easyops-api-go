package scheduler

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-scheduler/task"

	"github.com/easyopsapis/easyops-api-go/protorepo-scheduler/task_history"
)

type Client struct {
	Task task.Client

	TaskHistory task_history.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Task = task.NewClient(c)

	client.TaskHistory = task_history.NewClient(c)
	return &client
}
