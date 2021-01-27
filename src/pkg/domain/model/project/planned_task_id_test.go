// This file is auto generated

package project

import (
	"reflect"
	"testing"
)

func TestNewPlannedTaskID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    PlannedTaskID
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{id: "test-id"},
			want:    PlannedTaskID{id: "test-id"},
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args{id: ""},
			want:    PlannedTaskIDZero,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlannedTaskID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlannedTaskID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("NewPlannedTaskID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustNewPlannedTaskID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		want      PlannedTaskID
		wantPanic bool
	}{
		{
			name:      "ok",
			args:      args{id: "test-id"},
			want:      PlannedTaskID{id: "test-id"},
			wantPanic: false,
		},
		{
			name:      "ng",
			args:      args{id: ""},
			want:      PlannedTaskIDZero,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("MustNewPlannedTaskID() error = %v, wantPanic %v", err, tt.wantPanic)
				}
			}()
			got := MustNewPlannedTaskID(tt.args.id)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("MustNewPlannedTaskID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlannedTaskID_Equals(t *testing.T) {
	tests := []struct {
		name string
		id1  PlannedTaskID
		id2  PlannedTaskID
		want bool
	}{
		{
			name: "true",
			id1:  PlannedTaskID{id: "id1"},
			id2:  PlannedTaskID{id: "id2"},
			want: false,
		},
		{
			name: "true",
			id1:  PlannedTaskID{id: "id1"},
			id2:  PlannedTaskID{id: "id1"},
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

func TestPlannedTaskID_String(t *testing.T) {
	id, _ := NewPlannedTaskID("test-id")
	var (
		got  = id.String()
		want = "test-id"
	)
	if got != want {
		t.Errorf("PlannedTaskID.String() = %v, want %v", got, want)
	}
}
