package micro_app

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-micro_app/installed_micro_app"
)

type Client struct {
	InstalledMicroApp installed_micro_app.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.InstalledMicroApp = installed_micro_app.NewClient(c)
	return &client
}
