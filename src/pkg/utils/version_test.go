package utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestVersionMarshal(t *testing.T) {
	tests := []struct {
		name    string
		v       interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "naked",
			v:       NewVersion(1, 2, 3),
			want:    `"1.2.3"`,
			wantErr: false,
		},
		{
			name:    "with-suffix string",
			v:       Version{1, 2, 3, "suffix"},
			want:    `"1.2.3-suffix"`,
			wantErr: false,
		},
		{
			name: "in other struct",
			v: struct {
				Field1 string  `json:"field1"`
				Field2 Version `json:"field2"`
				Field3 int     `json:"field3"`
			}{
				Field1: "string value",
				Field2: Version{1, 2, 3, "suffix"},
				Field3: 1234567890,
			},
			want:    `{"field1":"string value","field2":"1.2.3-suffix","field3":1234567890}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.v)
			if err != nil {
				t.Errorf("json.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("json.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersionUnmarshal(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		got     func(b []byte) (interface{}, error)
		want    interface{}
		wantErr bool
	}{
		{
			name: "naked",
			json: `"1.2.3-suffix"`,
			got: func(b []byte) (interface{}, error) {
				var v Version
				if err := json.Unmarshal(b, &v); err != nil {
					return nil, err
				}
				return v, nil
			},
			want:    Version{1, 2, 3, "suffix"},
			wantErr: false,
		},
		{
			name: "in other struct",
			json: `{"string_field":"1.2.3-suffix","version_field":"1.2.3-suffix"}`,
			got: func(b []byte) (interface{}, error) {
				var v struct {
					F1 string  `json:"string_field"`
					F2 Version `json:"version_field"`
				}
				if err := json.Unmarshal(b, &v); err != nil {
					return nil, err
				}
				return v, nil
			},
			want: struct {
				F1 string  `json:"string_field"`
				F2 Version `json:"version_field"`
			}{
				F1: "1.2.3-suffix",
				F2: Version{1, 2, 3, "suffix"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.got([]byte(tt.json))
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v, wantErr false", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshl() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
