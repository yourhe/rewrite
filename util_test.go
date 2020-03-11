package rewrite

import (
	"testing"
)

func TestDecodeHost(t *testing.T) {

	tests := []struct {
		name string
		host string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			host: "port-89.www.baidu.com",
			want: "www.baidu.com:89",
		},
		{
			name: "1",
			host: "www.baidu.com.port-89",
			want: "www.baidu.com:89",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeHost(tt.host, "port-", "."); got != tt.want {
				t.Errorf("DecodeHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeHost(t *testing.T) {

	tests := []struct {
		name string
		host string
		port string
		want string
	}{
		// TODO: Add test cases.
		{
			host: "www.baidu.com",
			want: "www.baidu.com",
		},
		{
			host: "www.baidu.com:99",
			want: "port-99.www.baidu.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeHost(tt.host, PortPrefix, "."); got != tt.want {
				t.Errorf("EncodeHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSWLetDigHyp(t *testing.T) {
	type args struct {
		host     string
		protocol string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				host:     "www.wanfangdata.com.cn",
				protocol: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSWLetDigHyp(tt.args.host, tt.args.protocol); got != tt.want {
				t.Errorf("ToSWLetDigHyp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSWReduceLetDigHyp(t *testing.T) {
	type args struct {
		host     string
		protocol string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				host:     "www.wanfangdata.com.cn",
				protocol: "",
			},
			want: "cn/com/wanfangdata/www",
		},
		{
			args: args{
				host:     "wanfangdata.com.cn",
				protocol: "",
			},
			want: "cn/com/wanfangdata",
		},
		{
			args: args{
				host:     ".wanfangdata.com.cn",
				protocol: "",
			},
			want: "cn/com/wanfangdata",
		},
		{
			args: args{
				host:     "...wanfangdata.com.cn",
				protocol: "",
			},
			want: "cn/com/wanfangdata",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSWReduceLetDigHyp(tt.args.host, tt.args.protocol); got != tt.want {
				t.Errorf("ToSWReduceLetDigHyp() = %v, want %v", got, tt.want)
			}
		})
	}
}
