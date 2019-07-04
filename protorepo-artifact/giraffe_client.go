package artifact

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-artifact/pkg"

	"github.com/easyopsapis/easyops-api-go/protorepo-artifact/version"
)

type Client struct {
	Pkg pkg.Client

	Version version.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Pkg = pkg.NewClient(c)

	client.Version = version.NewClient(c)
	return &client
}
