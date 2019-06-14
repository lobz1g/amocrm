package models

import (
	"reflect"
	"testing"
)

func TestAcc_Get(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name    string
		fields  fields
		want    *account
		wantErr bool
	}{
		{"error", fields{request{}}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			a := Acc{
				request: tt.fields.request,
			}
			got, err := a.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Acc.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Acc.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
