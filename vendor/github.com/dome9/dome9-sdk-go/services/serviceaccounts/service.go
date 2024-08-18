package serviceaccounts

import (
	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/dome9/client"
)

type Service struct {
	Client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{Client: client.NewClient(c)}
}
