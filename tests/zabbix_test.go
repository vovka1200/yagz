package tests

import (
	"reflect"
	"testing"

	zabbix "git.grav.su/andromeda/yagz"
)

func TestNewClient(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want zabbix.Zabbix
	}{
		{
			name: "test1",
			want: zabbix.Zabbix{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zabbix.NewClient(tt.args.url)
			tt.want.Client = got.Client
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient()\n got %#v\nwant %#v", got, tt.want)
			}
		})
	}
}
