package monitor

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/collector"
)

type Client struct {
	Collector collector.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Collector = collector.NewClient(c)
	return &client
}
