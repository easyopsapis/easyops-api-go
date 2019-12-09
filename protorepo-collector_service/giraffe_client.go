package collector_service

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-collector_service/jobs"
)

type Client struct {
	Jobs jobs.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Jobs = jobs.NewClient(c)
	return &client
}
