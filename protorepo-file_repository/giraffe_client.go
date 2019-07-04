package file_repository

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-file_repository/archive"

	"github.com/easyopsapis/easyops-api-go/protorepo-file_repository/workspace"
)

type Client struct {
	Archive archive.Client

	Workspace workspace.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Archive = archive.NewClient(c)

	client.Workspace = workspace.NewClient(c)
	return &client
}
