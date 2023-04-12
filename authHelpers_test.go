package goutils

import (
	"net/http"
	"testing"
)

func TestGetAuthHeader(t *testing.T) {

	type args struct {
		r *http.Request
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Authorization header only",
			args: args{
				r: &http.Request{
					Header: map[string][]string{
						"Authorization": {"Bearer 123"},
					},
				},
			},
			want: "Bearer 123",
		},
		{
			name: "X-Authorization header only",
			args: args{
				r: &http.Request{
					Header: map[string][]string{
						"X-Authorization": {"Bearer 123"},
					},
				},
			},
			want: "Bearer 123",
		},
		{
			name: "Authorization header and X-Authorization Header",
			args: args{
				r: &http.Request{
					Header: map[string][]string{
						"Authorization":   {"Bearer 123"},
						"X-Authorization": {"Bearer 123"},
					},
				},
			},
			want: "Bearer 123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAuthHeader(tt.args.r); got != tt.want {
				t.Errorf("GetAuthHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
