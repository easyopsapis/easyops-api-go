package notify

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/oplog"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/subscriber"
)

type Client struct {
	Oplog oplog.Client

	Subscriber subscriber.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Oplog = oplog.NewClient(c)

	client.Subscriber = subscriber.NewClient(c)
	return &client
}
