package test_test

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-test_test/hh"
)

type Client struct {
	Hh hh.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Hh = hh.NewClient(c)
	return &client
}
