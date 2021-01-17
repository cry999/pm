package web

import (
	"testing"
	"time"
)

func Test_time_Format(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want string
	}{
		{
			name: "JST(+0900)",
			date: time.Date(2020, 1, 23, 4, 56, 7, 8, time.FixedZone("JST", int(9*time.Hour/time.Second))),
			want: "2020-01-23T04:56:07+09:00",
		},
		{
			name: "UTC",
			date: time.Date(2020, 1, 23, 4, 56, 7, 8, time.UTC),
			want: "2020-01-23T04:56:07Z",
		},
		{
			name: "ETC(-0500)",
			date: time.Date(2020, 1, 23, 4, 56, 7, 8, time.FixedZone("EST", int(-5*time.Hour/time.Second))),
			want: "2020-01-23T04:56:07-05:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Format(time.RFC3339); got != tt.want {
				t.Errorf("time.Format = %v, want %v", got, tt.want)
			}
		})
	}
}
