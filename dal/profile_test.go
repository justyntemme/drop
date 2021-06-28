package dal

import (
	"context"
	"fmt"
	"testing"

	"gitlab.com/nextwavedevs/drop/models"
)

var dbc = DBContainer{
	Image: "mongo:4.4.5",
	Port:  "27017",
}

func TestGetUserById(t *testing.T) {
	log, db, _ := NewUnit(t, dbc)

	user := New(log, db)
	c := context.Background()
	type args struct {
		ctx     context.Context
		traceID string
		uid     string
	}
	tests := []struct {
		name string
		args args
		want models.User
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx:     c,
				traceID: "00000000-0000-0000-0000-000000000000",
				uid:     "60d8ce9b4f95827ce9961432",
			},

			want: models.User{
				UID:  "60d8ce9b4f95827ce9961432",
				Name: "justyn",
				City: "St Louis",
				Age:  24,
			},
		},
		{
			name: "empty string uid",
			args: args{
				ctx:     c,
				traceID: "00000000-0000-0000-0000-000000000000",
				uid:     "",
			},

			want: models.User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := user.GetUserById(tt.args.ctx, tt.args.traceID, tt.args.uid)
			if err != nil {
				fmt.Println("error in getuserbyid test")
			}
			if got != tt.want {
				t.Errorf("GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_updateProfile(t *testing.T) {
// 	log, db, teardown := NewUnit(t, dbc)
// 	t.Cleanup(teardown)

// 	user := New(log, db)

// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			user.UpdateProfile(tt.args.w, tt.args.r)
// 		})
// 	}
// }
