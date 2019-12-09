package micro_app

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-micro_app/installed_micro_app"

	"github.com/easyopsapis/easyops-api-go/protorepo-micro_app/micro_app_container"

	"github.com/easyopsapis/easyops-api-go/protorepo-micro_app/object_micro_app"

	"github.com/easyopsapis/easyops-api-go/protorepo-micro_app/report"
)

type Client struct {
	InstalledMicroApp installed_micro_app.Client

	MicroAppContainer micro_app_container.Client

	ObjectMicroApp object_micro_app.Client

	Report report.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.InstalledMicroApp = installed_micro_app.NewClient(c)

	client.MicroAppContainer = micro_app_container.NewClient(c)

	client.ObjectMicroApp = object_micro_app.NewClient(c)

	client.Report = report.NewClient(c)
	return &client
}
