package permission

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-permission/permission"

	"github.com/easyopsapis/easyops-api-go/protorepo-permission/role"
)

type Client struct {
	Permission permission.Client

	Role role.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Permission = permission.NewClient(c)

	client.Role = role.NewClient(c)
	return &client
}
