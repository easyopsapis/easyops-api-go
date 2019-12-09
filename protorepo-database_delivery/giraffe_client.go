package database_delivery

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-database_delivery/dbclient"

	"github.com/easyopsapis/easyops-api-go/protorepo-database_delivery/dbinstance"

	"github.com/easyopsapis/easyops-api-go/protorepo-database_delivery/dbservice"

	"github.com/easyopsapis/easyops-api-go/protorepo-database_delivery/dbtask"

	"github.com/easyopsapis/easyops-api-go/protorepo-database_delivery/sqlpkg_versions"

	"github.com/easyopsapis/easyops-api-go/protorepo-database_delivery/sqlpkgs"
)

type Client struct {
	Dbclient dbclient.Client

	Dbinstance dbinstance.Client

	Dbservice dbservice.Client

	Dbtask dbtask.Client

	SqlpkgVersions sqlpkg_versions.Client

	Sqlpkgs sqlpkgs.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Dbclient = dbclient.NewClient(c)

	client.Dbinstance = dbinstance.NewClient(c)

	client.Dbservice = dbservice.NewClient(c)

	client.Dbtask = dbtask.NewClient(c)

	client.SqlpkgVersions = sqlpkg_versions.NewClient(c)

	client.Sqlpkgs = sqlpkgs.NewClient(c)
	return &client
}
