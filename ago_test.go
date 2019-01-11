package ago

import (
	"testing"
	"time"
)

func TestAgo(t *testing.T) {

	now := time.Now().UTC()

	tests := []struct {
		input    string
		expected time.Time
	}{
		{"5 min ago", now.Add(-5 * time.Minute)},
		{"10 hours, 7 min ago", now.Add(-10 * time.Hour).Add(-7 * time.Minute)},
		{"24 days ago", now.Add(-24 * 24 * time.Hour)},
	}

	for _, test := range tests {
		v, err := Parse(test.input)
		if err != nil {
			t.Error(err)
		}
		if v.Unix() != test.expected.Unix() {
			t.Errorf("expected: %v, got %v", test.expected, v)
		}
	}

}
