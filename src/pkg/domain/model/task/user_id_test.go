// This file is auto generated

package task

import (
	"reflect"
	"testing"
)

func TestNewUserID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    UserID
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{id: "test-id"},
			want:    UserID{id: "test-id"},
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args{id: ""},
			want:    UserIDZero,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("NewUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustNewUserID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		want      UserID
		wantPanic bool
	}{
		{
			name:      "ok",
			args:      args{id: "test-id"},
			want:      UserID{id: "test-id"},
			wantPanic: false,
		},
		{
			name:      "ng",
			args:      args{id: ""},
			want:      UserIDZero,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("MustNewUserID() error = %v, wantPanic %v", err, tt.wantPanic)
				}
			}()
			got := MustNewUserID(tt.args.id)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("MustNewUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserID_Equals(t *testing.T) {
	tests := []struct {
		name string
		id1  UserID
		id2  UserID
		want bool
	}{
		{
			name: "true",
			id1:  UserID{id: "id1"},
			id2:  UserID{id: "id2"},
			want: false,
		},
		{
			name: "true",
			id1:  UserID{id: "id1"},
			id2:  UserID{id: "id1"},
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

func TestUserID_String(t *testing.T) {
	id, _ := NewUserID("test-id")
	var (
		got  = id.String()
		want = "test-id"
	)
	if got != want {
		t.Errorf("UserID.String() = %v, want %v", got, want)
	}
}
