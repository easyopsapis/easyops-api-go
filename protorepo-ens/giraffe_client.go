package ens

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-ens/legacy"
)

type Client struct {
	Legacy legacy.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Legacy = legacy.NewClient(c)
	return &client
}
