package main

import "testing"

func TestConditioner(t *testing.T) {
	tests := []struct {
		name     string
		tRoom    int
		tCond    int
		workMode string
		want     int
	}{
		{
			name:     "1",
			tRoom:    10,
			tCond:    20,
			workMode: "heat",
			want:     20,
		},
		{
			name:     "2",
			tRoom:    10,
			tCond:    20,
			workMode: "freeze",
			want:     10,
		},
		{
			name:     "3",
			tRoom:    0,
			tCond:    0,
			workMode: "freeze",
			want:     0,
		},
		{
			name:     "4",
			tRoom:    -10,
			tCond:    20,
			workMode: "auto",
			want:     20,
		},
		{
			name:     "5",
			tRoom:    10,
			tCond:    20,
			workMode: "fan",
			want:     10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Сonditioner(tt.tRoom, tt.tCond, tt.workMode)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
