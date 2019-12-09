package webshell

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-webshell/task"
)

type Client struct {
	Task task.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Task = task.NewClient(c)
	return &client
}
