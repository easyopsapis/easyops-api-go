package collector_center

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-collector_center/clazz"

	"github.com/easyopsapis/easyops-api-go/protorepo-collector_center/collection_config"

	"github.com/easyopsapis/easyops-api-go/protorepo-collector_center/configuration"

	"github.com/easyopsapis/easyops-api-go/protorepo-collector_center/job"

	"github.com/easyopsapis/easyops-api-go/protorepo-collector_center/template"
)

type Client struct {
	Clazz clazz.Client

	CollectionConfig collection_config.Client

	Configuration configuration.Client

	Job job.Client

	Template template.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Clazz = clazz.NewClient(c)

	client.CollectionConfig = collection_config.NewClient(c)

	client.Configuration = configuration.NewClient(c)

	client.Job = job.NewClient(c)

	client.Template = template.NewClient(c)
	return &client
}
