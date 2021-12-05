package dal

import (
	"context"
	"reflect"
	"testing"

	models "gitlab.com/nextwavedevs/drop/shared/models"
)

func TestGetListingById(t *testing.T) {
	type args struct {
		ctx     context.Context
		traceID string
		uid     string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Listing
		wantErr bool
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetListingById(tt.args.ctx, tt.args.traceID, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListingById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListingById() = %v, want %v", got, tt.want)
			}
		})
	}
}
