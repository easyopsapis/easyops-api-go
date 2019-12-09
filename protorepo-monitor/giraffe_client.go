package monitor

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/alert"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/alert_rule"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/app_health"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/collector"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/data_name"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/influxdb"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/log_search"

	"github.com/easyopsapis/easyops-api-go/protorepo-monitor/translate"
)

type Client struct {
	Alert alert.Client

	AlertRule alert_rule.Client

	AppHealth app_health.Client

	Collector collector.Client

	DataName data_name.Client

	Influxdb influxdb.Client

	LogSearch log_search.Client

	Translate translate.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Alert = alert.NewClient(c)

	client.AlertRule = alert_rule.NewClient(c)

	client.AppHealth = app_health.NewClient(c)

	client.Collector = collector.NewClient(c)

	client.DataName = data_name.NewClient(c)

	client.Influxdb = influxdb.NewClient(c)

	client.LogSearch = log_search.NewClient(c)

	client.Translate = translate.NewClient(c)
	return &client
}
