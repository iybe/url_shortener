package services

import "testing"

func TestGenerateStringRandom(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name     string
		args     args
		wantSize int
	}{
		{
			name: "success",
			args: args{
				size: 6,
			},
			wantSize: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateStringRandom(tt.args.size); len(got) != tt.wantSize {
				t.Errorf("GenerateStringRandom() = %d, want %d", len(got), tt.wantSize)
			}
		})
	}
}
