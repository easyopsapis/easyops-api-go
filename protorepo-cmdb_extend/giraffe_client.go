package cmdb_extend

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/agent"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/agent_install_script"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/custom"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/email"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/host_resolved"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/instance"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/org"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/pipeline"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb_extend/toolkit"
)

type Client struct {
	Agent agent.Client

	AgentInstallScript agent_install_script.Client

	Custom custom.Client

	Email email.Client

	HostResolved host_resolved.Client

	Instance instance.Client

	Org org.Client

	Pipeline pipeline.Client

	Toolkit toolkit.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Agent = agent.NewClient(c)

	client.AgentInstallScript = agent_install_script.NewClient(c)

	client.Custom = custom.NewClient(c)

	client.Email = email.NewClient(c)

	client.HostResolved = host_resolved.NewClient(c)

	client.Instance = instance.NewClient(c)

	client.Org = org.NewClient(c)

	client.Pipeline = pipeline.NewClient(c)

	client.Toolkit = toolkit.NewClient(c)
	return &client
}
