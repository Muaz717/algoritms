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

func TestPhoneNumbers(t *testing.T) {
	tests := []struct {
		name      string
		phoneNums []string
		want      string
	}{
		{
			name:      "1",
			phoneNums: []string{"8(495)430-23-97", "+7-4-9-5-43-023-97", "4-3-0-2-3-9-7", "8-495-430"},
			want:      "YES, YES, NO",
		},
		{
			name:      "2",
			phoneNums: []string{"8(928)430-23-97", "+7-9-2-8-43-023-97", "4-3-0-2-3-9-7", "8-495-430"},
			want:      "YES, NO, NO",
		},
		{
			name:      "3",
			phoneNums: []string{"4-3-0-2-3-9-7", "8-495-430", "8(495)430-23-97", "+7-4-9-5-43-023-97"},
			want:      "NO, YES, YES",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PhoneNumbers(tt.phoneNums[0], tt.phoneNums[1], tt.phoneNums[2], tt.phoneNums[3])
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
