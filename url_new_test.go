package rewrite

import (
	"reflect"
	"testing"
)

func Test_recodePath(t *testing.T) {
	tests := []struct {
		name string
		path []byte

		want []byte
	}{
		// TODO: Add test cases.
		{
			path: []byte("a.html"),
			want: []byte("/a.html"),
		},
		{
			path: []byte("/a.html"),
			want: []byte("/a.html"),
		},
		{
			path: []byte("/../a.html"),
			want: []byte("/a.html"),
		},
		{
			path: []byte("/../../../a.html"),
			want: []byte("/a.html"),
		},
		{
			path: []byte("/p1/g/../a.html"),
			want: []byte("/p1/a.html"),
		},
		{
			path: []byte("/p1/g/../../a.html"),
			want: []byte("/a.html"),
		},
		{
			path: []byte("/p1/g/../cc/../a.html"),
			want: []byte("/p1/a.html"),
		},
		{
			path: []byte("/p1/g/../cc/../../a.html"),
			want: []byte("/a.html"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recodePath(tt.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recodePath() = %s, want %s", got, tt.want)
			}
		})
	}
}
