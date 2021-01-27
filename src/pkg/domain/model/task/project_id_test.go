// This file is auto generated

package task

import (
	"reflect"
	"testing"
)

func TestNewProjectID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    ProjectID
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{id: "test-id"},
			want:    ProjectID{id: "test-id"},
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args{id: ""},
			want:    ProjectIDZero,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProjectID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProjectID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("NewProjectID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustNewProjectID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		want      ProjectID
		wantPanic bool
	}{
		{
			name:      "ok",
			args:      args{id: "test-id"},
			want:      ProjectID{id: "test-id"},
			wantPanic: false,
		},
		{
			name:      "ng",
			args:      args{id: ""},
			want:      ProjectIDZero,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("MustNewProjectID() error = %v, wantPanic %v", err, tt.wantPanic)
				}
			}()
			got := MustNewProjectID(tt.args.id)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("MustNewProjectID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectID_Equals(t *testing.T) {
	tests := []struct {
		name string
		id1  ProjectID
		id2  ProjectID
		want bool
	}{
		{
			name: "true",
			id1:  ProjectID{id: "id1"},
			id2:  ProjectID{id: "id2"},
			want: false,
		},
		{
			name: "true",
			id1:  ProjectID{id: "id1"},
			id2:  ProjectID{id: "id1"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.id1.Equals(tt.id2); got != tt.want {
				t.Errorf("'%s'.Equlas('%s') = %v, want %v", tt.id1, tt.id2, got, tt.want)
			}
		})
	}
}

func TestProjectID_String(t *testing.T) {
	id, _ := NewProjectID("test-id")
	var (
		got  = id.String()
		want = "test-id"
	)
	if got != want {
		t.Errorf("ProjectID.String() = %v, want %v", got, want)
	}
}
