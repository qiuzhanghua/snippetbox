package main

import (
	"taiji.dev/snippetbox/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC)
	hd := humanDate(tm)
	if hd != "17 Mar 2022 at 10:15" {
		t.Errorf("want %q; got %q", "17 Mar 2022 at 10:15", hd)
	}

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2021, 2, 17, 10, 0, 0, 0, time.UTC),
			want: "17 Feb 2021 at 10:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2021, 2, 17, 10, 0, 0, 0, time.FixedZone("CET", 60*60)),
			want: "17 Feb 2021 at 09:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}
