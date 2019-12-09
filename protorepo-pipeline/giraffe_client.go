package pipeline

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-pipeline/build"

	"github.com/easyopsapis/easyops-api-go/protorepo-pipeline/org"

	"github.com/easyopsapis/easyops-api-go/protorepo-pipeline/pipeline"

	"github.com/easyopsapis/easyops-api-go/protorepo-pipeline/project"

	"github.com/easyopsapis/easyops-api-go/protorepo-pipeline/provider"

	"github.com/easyopsapis/easyops-api-go/protorepo-pipeline/template"
)

type Client struct {
	Build build.Client

	Org org.Client

	Pipeline pipeline.Client

	Project project.Client

	Provider provider.Client

	Template template.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Build = build.NewClient(c)

	client.Org = org.NewClient(c)

	client.Pipeline = pipeline.NewClient(c)

	client.Project = project.NewClient(c)

	client.Provider = provider.NewClient(c)

	client.Template = template.NewClient(c)
	return &client
}
