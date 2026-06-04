package tests

import (
	"context"
	"testing"

	"git.grav.su/andromeda/yagz"
	"git.grav.su/andromeda/yagz/generic"
	"git.grav.su/andromeda/yagz/hostgroup"
	"git.grav.su/andromeda/yagz/hostgroup/get"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	var err error
	log.SetLevel(log.DebugLevel)
	client := zabbix.NewClient("http://172.22.231.9:8080/api_jsonrpc.php")
	err = client.Login("Admin", "xhoK4Dowb3~HBX6")
	assert.NoError(t, err)

	var hostGroups hostgroup.HostGroups
	ctx := context.Background()
	err = client.Call(ctx, &hostGroups, get.Method, get.Params{
		Params: generic.Params{
			Filter: generic.Filter{
				"name": "FVF",
			},
		},
	})
	assert.NoError(t, err)
	t.Log(hostGroups)
}
