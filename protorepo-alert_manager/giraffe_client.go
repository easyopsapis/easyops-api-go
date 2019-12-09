package alert_manager

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-alert_manager/rule_list"
)

type Client struct {
	RuleList rule_list.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.RuleList = rule_list.NewClient(c)
	return &client
}
