package black_jack

import "testing"

func Test_CalculatePoints(t *testing.T) {
	type args struct {
		cards    []string
		expected string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "2, 3 = 5",
			args: args{
				cards:    []string{"2", "3"},
				expected: "5",
			},
		},
		// ...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ...
			// assert.Equal(t, ..., ...)
		})
	}
}
