package dal

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserById(t *testing.T) {
	c := context.Background()
	writer := httptest.NewRecorder()
	type args struct {
		ctx context.Context
		w   http.ResponseWriter
		uid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx: c,
				w:   writer,
				uid: "60035d152f2355126396353d",
			},

			want: "[{\"City\":\"St. Louis\",\"ID\":\"00001\",\"_id\":\"60035d152f2355126396353d\",\"age\":24,\"name\":\"justyn\",\"uid\":\"60035d152f2355126396353d\"}]",
		},
		{
			name: "empty string uid",
			args: args{
				ctx: c,
				w:   writer,
				uid: "",
			},

			want: "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserById(tt.args.ctx, tt.args.w, tt.args.uid); got != tt.want {
				t.Errorf("GetUserById() = %v, want %v", got, tt.want)

			}
		})
	}
}

func Test_updateProfile(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateProfile(tt.args.w, tt.args.r)
		})
	}
}
