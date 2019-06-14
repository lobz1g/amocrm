package models

import (
	"reflect"
	"testing"
)

func TestNt_Create(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name   string
		fields fields
		want   *note
	}{
		{"", fields{request{}}, &note{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Nt{
				request: tt.fields.request,
			}
			if got := n.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Nt.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNt_Add(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		nt *note
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"error", fields{request{}}, args{&note{}}, 0, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			n := Nt{
				request: tt.fields.request,
			}
			got, err := n.Add(tt.args.nt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Nt.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Nt.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
