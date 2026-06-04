package zabbix

import (
	"context"

	"git.grav.su/andromeda/jsonrpc"
	"git.grav.su/andromeda/yagz/user/login"
	log "github.com/sirupsen/logrus"
)

type Zabbix struct {
	Client  jsonrpc.RPCClient
	Version string
	User    login.User
}

func NewClient(url string) Zabbix {
	return Zabbix{
		Client: jsonrpc.NewClientWithOpts(url, &jsonrpc.RPCClientOpts{
			AllowUnknownFields: true,
			DefaultRequestID:   1,
			CustomHeaders: map[string]string{
				"User-Agent": "yagz/0.0.1",
			},
		}),
	}
}

func NewClientWithOpts(url string, options *jsonrpc.RPCClientOpts) Zabbix {
	options.AllowUnknownFields = true
	return Zabbix{
		Client: jsonrpc.NewClientWithOpts(url, options),
	}
}

func (z *Zabbix) Login(username string, password string) error {
	ctx := context.Background()
	var token string
	if err := z.Call(ctx, &token, login.Method, login.Login{
		Username: username,
		Password: password,
	}); err == nil {
		z.Client.SetToken(token)
		return nil
	} else {
		log.Error(err)
		return err
	}
}

func (z *Zabbix) Call(ctx context.Context, result any, method string, params any) error {
	return z.Client.CallFor(ctx, result, method, params)
}
