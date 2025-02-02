package zabbix

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/vovka1200/jsonrpc"
	"github.com/vovka1200/yagz/user/login"
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
	var user login.User
	if err := z.Call(ctx, &user, login.Method, login.Login{
		User:     username,
		Password: password,
		UserData: true,
	}); err == nil {
		z.User = user
		z.Client.SetToken(z.User.SessionId)
		return nil
	} else {
		log.Error(err)
		return err
	}
}

func (z *Zabbix) Call(ctx context.Context, result any, method string, params any) error {
	return z.Client.CallFor(ctx, result, method, params)
}
