package models

import (
	"reflect"
	"testing"
)

func TestCmpn_Create(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name   string
		fields fields
		want   *company
	}{
		{"", fields{request{}}, &company{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cmpn{
				request: tt.fields.request,
			}
			if got := c.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmpn.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmpn_All(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*company
		wantErr bool
	}{
		{"error", fields{request{}}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			c := Cmpn{
				request: tt.fields.request,
			}
			got, err := c.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cmpn.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmpn.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmpn_Responsible(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*company
		wantErr bool
	}{
		{"error", fields{request{}}, args{0}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			c := Cmpn{
				request: tt.fields.request,
			}
			got, err := c.Responsible(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cmpn.Responsible() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmpn.Responsible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmpn_Id(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *company
		wantErr bool
	}{
		{"error", fields{request{}}, args{0}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			c := Cmpn{
				request: tt.fields.request,
			}
			got, err := c.Id(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cmpn.Id() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmpn.Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmpn_Add(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		cmpn *company
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"error", fields{request{}}, args{&company{}}, 0, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			c := Cmpn{
				request: tt.fields.request,
			}
			got, err := c.Add(tt.args.cmpn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cmpn.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cmpn.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmpn_Update(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		cmpn *company
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"error", fields{request{}}, args{&company{}}, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			c := Cmpn{
				request: tt.fields.request,
			}
			if err := c.Update(tt.args.cmpn); (err != nil) != tt.wantErr {
				t.Errorf("Cmpn.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
