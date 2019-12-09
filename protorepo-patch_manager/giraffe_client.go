package patch_manager

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-patch_manager/patch"
)

type Client struct {
	Patch patch.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Patch = patch.NewClient(c)
	return &client
}
