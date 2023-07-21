package models

import (
	"taiji.dev/snippetbox/internal/assert"
	"testing"
)

func TestUserModelExists(t *testing.T) {
	// Set up a suite of table-driven tests and expected results.
	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Non-existent ID",
			userID: 2,
			want:   false,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)
			m := UserModel{db}
			exists, err := m.Exists(tt.userID)
			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}
