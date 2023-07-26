package utils

import "testing"

func TestIsDirectlyRun(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "t1",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDirectlyRun(); got != tt.want {
				t.Errorf("IsDirectlyRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
