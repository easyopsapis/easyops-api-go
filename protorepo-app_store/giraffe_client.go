package app_store

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-app_store/micro_app"

	"github.com/easyopsapis/easyops-api-go/protorepo-app_store/version"
)

type Client struct {
	MicroApp micro_app.Client

	Version version.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.MicroApp = micro_app.NewClient(c)

	client.Version = version.NewClient(c)
	return &client
}
