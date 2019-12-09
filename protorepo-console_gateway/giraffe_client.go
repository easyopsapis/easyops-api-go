package console_gateway

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-console_gateway/cmdb_service_ctrl"

	"github.com/easyopsapis/easyops-api-go/protorepo-console_gateway/notify_ctrl"

	"github.com/easyopsapis/easyops-api-go/protorepo-console_gateway/permission_ctrl"

	"github.com/easyopsapis/easyops-api-go/protorepo-console_gateway/user_service_ctrl"
)

type Client struct {
	CmdbServiceCtrl cmdb_service_ctrl.Client

	NotifyCtrl notify_ctrl.Client

	PermissionCtrl permission_ctrl.Client

	UserServiceCtrl user_service_ctrl.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.CmdbServiceCtrl = cmdb_service_ctrl.NewClient(c)

	client.NotifyCtrl = notify_ctrl.NewClient(c)

	client.PermissionCtrl = permission_ctrl.NewClient(c)

	client.UserServiceCtrl = user_service_ctrl.NewClient(c)
	return &client
}
