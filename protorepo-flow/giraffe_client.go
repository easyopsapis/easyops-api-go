package flow

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-flow/basic"

	"github.com/easyopsapis/easyops-api-go/protorepo-flow/execute"
)

type Client struct {
	Basic basic.Client

	Execute execute.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Basic = basic.NewClient(c)

	client.Execute = execute.NewClient(c)
	return &client
}
