package amocrm

import (
	"reflect"
	"testing"
)

func TestNewAmo(t *testing.T) {
	type args struct {
		login  string
		key    string
		domain string
	}
	tests := []struct {
		name string
		args args
		want *amo
	}{
		{"", args{"", "", ""}, &amo{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAmo(tt.args.login, tt.args.key, tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAmo() = %v, want %v", got, tt.want)
			}
		})
	}
}
