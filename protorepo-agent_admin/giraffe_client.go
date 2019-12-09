package agent_admin

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-agent_admin/admin_task"

	"github.com/easyopsapis/easyops-api-go/protorepo-agent_admin/agent"

	"github.com/easyopsapis/easyops-api-go/protorepo-agent_admin/org_init"

	"github.com/easyopsapis/easyops-api-go/protorepo-agent_admin/plugin"

	"github.com/easyopsapis/easyops-api-go/protorepo-agent_admin/plugin_version"
)

type Client struct {
	AdminTask admin_task.Client

	Agent agent.Client

	OrgInit org_init.Client

	Plugin plugin.Client

	PluginVersion plugin_version.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.AdminTask = admin_task.NewClient(c)

	client.Agent = agent.NewClient(c)

	client.OrgInit = org_init.NewClient(c)

	client.Plugin = plugin.NewClient(c)

	client.PluginVersion = plugin_version.NewClient(c)
	return &client
}
