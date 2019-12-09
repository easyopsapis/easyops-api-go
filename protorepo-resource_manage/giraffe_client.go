package resource_manage

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-resource_manage/cmdb_approve"

	"github.com/easyopsapis/easyops-api-go/protorepo-resource_manage/collector_history"

	"github.com/easyopsapis/easyops-api-go/protorepo-resource_manage/service_manage"
)

type Client struct {
	CmdbApprove cmdb_approve.Client

	CollectorHistory collector_history.Client

	ServiceManage service_manage.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.CmdbApprove = cmdb_approve.NewClient(c)

	client.CollectorHistory = collector_history.NewClient(c)

	client.ServiceManage = service_manage.NewClient(c)
	return &client
}
