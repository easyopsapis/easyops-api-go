package notify

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/msgpub"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/oplog"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/org"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/subscriber"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/subscriber_manager"

	"github.com/easyopsapis/easyops-api-go/protorepo-notify/topic"
)

type Client struct {
	Msgpub msgpub.Client

	Oplog oplog.Client

	Org org.Client

	Subscriber subscriber.Client

	SubscriberManager subscriber_manager.Client

	Topic topic.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Msgpub = msgpub.NewClient(c)

	client.Oplog = oplog.NewClient(c)

	client.Org = org.NewClient(c)

	client.Subscriber = subscriber.NewClient(c)

	client.SubscriberManager = subscriber_manager.NewClient(c)

	client.Topic = topic.NewClient(c)
	return &client
}
