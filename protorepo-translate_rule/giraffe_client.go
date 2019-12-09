package translate_rule

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-translate_rule/date_name_list_handler"

	"github.com/easyopsapis/easyops-api-go/protorepo-translate_rule/translate_rule_list_handler"
)

type Client struct {
	DateNameListHandler date_name_list_handler.Client

	TranslateRuleListHandler translate_rule_list_handler.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.DateNameListHandler = date_name_list_handler.NewClient(c)

	client.TranslateRuleListHandler = translate_rule_list_handler.NewClient(c)
	return &client
}
