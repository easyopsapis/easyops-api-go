package metadata_center

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-metadata_center/stream"
)

type Client struct {
	Stream stream.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Stream = stream.NewClient(c)
	return &client
}
