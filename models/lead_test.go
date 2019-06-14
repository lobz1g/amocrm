package models

import (
	"reflect"
	"testing"
)

func TestLd_Create(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name   string
		fields fields
		want   *lead
	}{
		{"", fields{request{}}, &lead{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Ld{
				request: tt.fields.request,
			}
			if got := l.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ld.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLd_All(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*lead
		wantErr bool
	}{
		{"error", fields{request{}}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			l := Ld{
				request: tt.fields.request,
			}
			got, err := l.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ld.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ld.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLd_Responsible(t *testing.T) {
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
		want    []*lead
		wantErr bool
	}{
		{"error", fields{request{}}, args{0}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			l := Ld{
				request: tt.fields.request,
			}
			got, err := l.Responsible(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ld.Responsible() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ld.Responsible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLd_Status(t *testing.T) {
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
		want    []*lead
		wantErr bool
	}{
		{"error", fields{request{}}, args{0}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			l := Ld{
				request: tt.fields.request,
			}
			got, err := l.Status(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ld.Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ld.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLd_Id(t *testing.T) {
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
		want    *lead
		wantErr bool
	}{
		{"error", fields{request{}}, args{0}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			l := Ld{
				request: tt.fields.request,
			}
			got, err := l.Id(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ld.Id() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ld.Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLd_Add(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		ld *lead
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"error", fields{request{}}, args{&lead{}}, 0, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			l := Ld{
				request: tt.fields.request,
			}
			got, err := l.Add(tt.args.ld)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ld.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ld.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLd_Update(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		ld *lead
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"error", fields{request{}}, args{&lead{}}, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			l := Ld{
				request: tt.fields.request,
			}
			if err := l.Update(tt.args.ld); (err != nil) != tt.wantErr {
				t.Errorf("Ld.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
