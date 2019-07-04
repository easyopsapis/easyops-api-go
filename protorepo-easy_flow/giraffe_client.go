package easy_flow

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-easy_flow/deploy"

	"github.com/easyopsapis/easyops-api-go/protorepo-easy_flow/instance"

	"github.com/easyopsapis/easyops-api-go/protorepo-easy_flow/strategy"
)

type Client struct {
	Deploy deploy.Client

	Instance instance.Client

	Strategy strategy.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Deploy = deploy.NewClient(c)

	client.Instance = instance.NewClient(c)

	client.Strategy = strategy.NewClient(c)
	return &client
}
