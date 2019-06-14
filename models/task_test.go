package models

import (
	"reflect"
	"testing"
)

func TestTsk_Create(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name   string
		fields fields
		want   *task
	}{
		{"", fields{request{}}, &task{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpTask := Tsk{
				request: tt.fields.request,
			}
			if got := tmpTask.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tsk.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTsk_All(t *testing.T) {
	type fields struct {
		request request
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*task
		wantErr bool
	}{
		{"error", fields{request{}}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			tmpTask := Tsk{
				request: tt.fields.request,
			}
			got, err := tmpTask.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tsk.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tsk.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTsk_Id(t *testing.T) {
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
		want    *task
		wantErr bool
	}{
		{"error", fields{request{}}, args{0}, nil, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			tmpTask := Tsk{
				request: tt.fields.request,
			}
			got, err := tmpTask.Id(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tsk.Id() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tsk.Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTsk_Add(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		tsk *task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"error", fields{request{}}, args{&task{}}, 0, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			tmpTask := Tsk{
				request: tt.fields.request,
			}
			got, err := tmpTask.Add(tt.args.tsk)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tsk.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Tsk.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTsk_Update(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		tsk *task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"error", fields{request{}}, args{&task{}}, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			tmpTask := Tsk{
				request: tt.fields.request,
			}
			if err := tmpTask.Update(tt.args.tsk); (err != nil) != tt.wantErr {
				t.Errorf("Tsk.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTsk_Close(t *testing.T) {
	type fields struct {
		request request
	}
	type args struct {
		tsk *task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"error", fields{request{}}, args{&task{}}, true},
	}
	for _, tt := range tests {
		if tt.name == "error" {
			OpenConnection("error", "error", "error")
		}
		t.Run(tt.name, func(t *testing.T) {
			tmpTask := Tsk{
				request: tt.fields.request,
			}
			if err := tmpTask.Close(tt.args.tsk); (err != nil) != tt.wantErr {
				t.Errorf("Tsk.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
