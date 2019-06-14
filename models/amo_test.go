package models

import "testing"

func TestOpenConnection(t *testing.T) {
	type args struct {
		login  string
		key    string
		domain string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"error", args{"error", "error", "error"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OpenConnection(tt.args.login, tt.args.key, tt.args.domain); (err != nil) != tt.wantErr {
				t.Errorf("OpenConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
