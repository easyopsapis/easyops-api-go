package user_service

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-user_service/apikey"

	"github.com/easyopsapis/easyops-api-go/protorepo-user_service/auth"

	"github.com/easyopsapis/easyops-api-go/protorepo-user_service/gateway"

	"github.com/easyopsapis/easyops-api-go/protorepo-user_service/invitation_code"

	"github.com/easyopsapis/easyops-api-go/protorepo-user_service/organization"

	"github.com/easyopsapis/easyops-api-go/protorepo-user_service/user_admin"
)

type Client struct {
	Apikey apikey.Client

	Auth auth.Client

	Gateway gateway.Client

	InvitationCode invitation_code.Client

	Organization organization.Client

	UserAdmin user_admin.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Apikey = apikey.NewClient(c)

	client.Auth = auth.NewClient(c)

	client.Gateway = gateway.NewClient(c)

	client.InvitationCode = invitation_code.NewClient(c)

	client.Organization = organization.NewClient(c)

	client.UserAdmin = user_admin.NewClient(c)
	return &client
}
