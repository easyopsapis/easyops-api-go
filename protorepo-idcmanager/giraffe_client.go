package idcmanager

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-idcmanager/excel"

	"github.com/easyopsapis/easyops-api-go/protorepo-idcmanager/idc"

	"github.com/easyopsapis/easyops-api-go/protorepo-idcmanager/idcrack"

	"github.com/easyopsapis/easyops-api-go/protorepo-idcmanager/org"
)

type Client struct {
	Excel excel.Client

	Idc idc.Client

	Idcrack idcrack.Client

	Org org.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Excel = excel.NewClient(c)

	client.Idc = idc.NewClient(c)

	client.Idcrack = idcrack.NewClient(c)

	client.Org = org.NewClient(c)
	return &client
}
