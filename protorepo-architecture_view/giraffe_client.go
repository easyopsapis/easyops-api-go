package architecture_view

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-architecture_view/business"

	"github.com/easyopsapis/easyops-api-go/protorepo-architecture_view/org"
)

type Client struct {
	Business business.Client

	Org org.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Business = business.NewClient(c)

	client.Org = org.NewClient(c)
	return &client
}
