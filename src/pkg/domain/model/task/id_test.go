// This file is auto generated

package task

import (
	"reflect"
	"testing"
)

func TestNewID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    ID
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{id: "test-id"},
			want:    ID{id: "test-id"},
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args{id: ""},
			want:    IDZero,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("NewID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustNewID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		want      ID
		wantPanic bool
	}{
		{
			name:      "ok",
			args:      args{id: "test-id"},
			want:      ID{id: "test-id"},
			wantPanic: false,
		},
		{
			name:      "ng",
			args:      args{id: ""},
			want:      IDZero,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("MustNewID() error = %v, wantPanic %v", err, tt.wantPanic)
				}
			}()
			got := MustNewID(tt.args.id)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("MustNewID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_Equals(t *testing.T) {
	tests := []struct {
		name string
		id1  ID
		id2  ID
		want bool
	}{
		{
			name: "true",
			id1:  ID{id: "id1"},
			id2:  ID{id: "id2"},
			want: false,
		},
		{
			name: "true",
			id1:  ID{id: "id1"},
			id2:  ID{id: "id1"},
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

func TestID_String(t *testing.T) {
	id, _ := NewID("test-id")
	var (
		got  = id.String()
		want = "test-id"
	)
	if got != want {
		t.Errorf("ID.String() = %v, want %v", got, want)
	}
}
