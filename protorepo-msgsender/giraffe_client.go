package msgsender

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-msgsender/custom_sender"
)

type Client struct {
	CustomSender custom_sender.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.CustomSender = custom_sender.NewClient(c)
	return &client
}
