package ops_automation

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-ops_automation/job_task"

	"github.com/easyopsapis/easyops-api-go/protorepo-ops_automation/jobs"

	"github.com/easyopsapis/easyops-api-go/protorepo-ops_automation/menu"

	"github.com/easyopsapis/easyops-api-go/protorepo-ops_automation/org"

	"github.com/easyopsapis/easyops-api-go/protorepo-ops_automation/service_event"
)

type Client struct {
	JobTask job_task.Client

	Jobs jobs.Client

	Menu menu.Client

	Org org.Client

	ServiceEvent service_event.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.JobTask = job_task.NewClient(c)

	client.Jobs = jobs.NewClient(c)

	client.Menu = menu.NewClient(c)

	client.Org = org.NewClient(c)

	client.ServiceEvent = service_event.NewClient(c)
	return &client
}
