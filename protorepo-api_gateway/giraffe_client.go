package api_gateway

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-api_gateway/bootstrap"
)

type Client struct {
	Bootstrap bootstrap.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Bootstrap = bootstrap.NewClient(c)
	return &client
}
