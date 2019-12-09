package translate

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-translate/data_name"

	"github.com/easyopsapis/easyops-api-go/protorepo-translate/translate_rule"

	"github.com/easyopsapis/easyops-api-go/protorepo-translate/translate_schema"
)

type Client struct {
	DataName data_name.Client

	TranslateRule translate_rule.Client

	TranslateSchema translate_schema.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.DataName = data_name.NewClient(c)

	client.TranslateRule = translate_rule.NewClient(c)

	client.TranslateSchema = translate_schema.NewClient(c)
	return &client
}
