package inspection

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-inspection/collector"

	"github.com/easyopsapis/easyops-api-go/protorepo-inspection/history"

	"github.com/easyopsapis/easyops-api-go/protorepo-inspection/info"

	"github.com/easyopsapis/easyops-api-go/protorepo-inspection/metric_group"

	"github.com/easyopsapis/easyops-api-go/protorepo-inspection/task"

	"github.com/easyopsapis/easyops-api-go/protorepo-inspection/template"
)

type Client struct {
	Collector collector.Client

	History history.Client

	Info info.Client

	MetricGroup metric_group.Client

	Task task.Client

	Template template.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Collector = collector.NewClient(c)

	client.History = history.NewClient(c)

	client.Info = info.NewClient(c)

	client.MetricGroup = metric_group.NewClient(c)

	client.Task = task.NewClient(c)

	client.Template = template.NewClient(c)
	return &client
}
