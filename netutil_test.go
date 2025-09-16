package netutil

import "testing"

func TestParsePort(t *testing.T) {
	tests := []struct {
		in    string
		want  int
		valid bool
	}{
		{"80", 80, true},
		{" 443 ", 443, true},
		{"0", 0, false},
		{"70000", 0, false},
		{"abc", 0, false},
	}

	for _, tt := range tests {
		got, err := ParsePort(tt.in)
		if tt.valid && err != nil {
			t.Errorf("ParsePort(%q) unexpected error: %v", tt.in, err)
		}
		if !tt.valid && err == nil {
			t.Errorf("ParsePort(%q) expected error, got none", tt.in)
		}
		if got != tt.want {
			t.Errorf("ParsePort(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}
