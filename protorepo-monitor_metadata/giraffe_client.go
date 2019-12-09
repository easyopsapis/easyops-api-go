package monitor_metadata

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor_metadata/alert_rule"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor_metadata/translate"
)

type Client struct {
	AlertRule alert_rule.Client

	Translate translate.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.AlertRule = alert_rule.NewClient(c)

	client.Translate = translate.NewClient(c)
	return &client
}
