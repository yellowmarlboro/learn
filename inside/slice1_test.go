package inside

import "testing"

func Test_sliceL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceL()
		})
	}
}