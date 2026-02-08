package stats

import (
	"testing"
	"time"
)

func TestCalculateTypingSpeed(t *testing.T) {
	tests := []struct {
		name         string
		correctChars int
		duration     time.Duration
		want         float64
	}{
		{
			name:         "High typing speed",
			correctChars: 600,
			duration:     60 * time.Second,
			want:         120,
		},
		{
			name:         "Decent typing speed",
			correctChars: 400,
			duration:     60 * time.Second,
			want:         80,
		},
		{
			name:         "Aight typing speed",
			correctChars: 300,
			duration:     60 * time.Second,
			want:         60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTypingSpeed(tt.correctChars, tt.duration)
			if got != tt.want {
				t.Errorf("CalculateTypingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRawTypingSpeed(t *testing.T) {
	tests := []struct {
		name           string
		correctChars   int
		incorrectChars int
		duration       time.Duration
		want           float64
	}{
		{
			name:           "Good acc",
			correctChars:   600,
			incorrectChars: 19,
			duration:       60 * time.Second,
			want:           123.8,
		},
		{
			name:           "Decent typing acc",
			correctChars:   400,
			incorrectChars: 30,
			duration:       60 * time.Second,
			want:           86,
		},
		{
			name:           "Lot of errors",
			correctChars:   300,
			incorrectChars: 409,
			duration:       60 * time.Second,
			want:           141.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateRawTypingSpeed(tt.correctChars, tt.incorrectChars, tt.duration)
			if got != tt.want {
				t.Errorf("CalculateRawTypingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateAccuracy(t *testing.T) {
	tests := []struct {
		name           string
		correctChars   int
		allChars       int
		backspaceCount int
		want           float64
		wantErr        bool
	}{
		{
			name:           "100% accuracy",
			correctChars:   100,
			allChars:       100,
			backspaceCount: 0,
			want:           100,
			wantErr:        false,
		},
		{
			name:           "Error case: more correct than total",
			correctChars:   200,
			allChars:       100,
			backspaceCount: 0,
			want:           0,
			wantErr:        true,
		},
		{
			name:           "99% accuracy with backspaces",
			correctChars:   100,
			allChars:       100,
			backspaceCount: 1,
			want:           99,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateAccuracy(tt.correctChars, tt.allChars, tt.backspaceCount)

			// Check if error presence matches our expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAccuracy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If we didn't want an error, check the result
			if !tt.wantErr && got != tt.want {
				t.Errorf("CalculateAccuracy() got = %v, want %v", got, tt.want)
			}
		})
	}
}
