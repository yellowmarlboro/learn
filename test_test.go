package main

import "testing"

func Test_forTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test_one"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			forTest()
		})
	}
}
