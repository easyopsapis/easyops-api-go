package tool

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-tool/basic"

	"github.com/easyopsapis/easyops-api-go/protorepo-tool/execute"

	"github.com/easyopsapis/easyops-api-go/protorepo-tool/migrate"
)

type Client struct {
	Basic basic.Client

	Execute execute.Client

	Migrate migrate.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Basic = basic.NewClient(c)

	client.Execute = execute.NewClient(c)

	client.Migrate = migrate.NewClient(c)
	return &client
}
