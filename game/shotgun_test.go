package game

import (
	"testing"
)

func Test_shotgun_LiveShells(t *testing.T) {
	type fields struct {
		shellsLeft uint8
		chamber    uint8
		dmg        uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		{
			name: "4 shells, 2 live, 2 blanks",
			fields: fields{
				shellsLeft: 4,
				chamber:    0b11110110, // last 4 bits matter
			},
			want: 2,
		},
		{
			name: "0 shells left",
			fields: fields{
				shellsLeft: 0,
				chamber:    0x64,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shotgun{
				shellsLeft: tt.fields.shellsLeft,
				chamber:    tt.fields.chamber,
				dmg:        tt.fields.dmg,
			}
			if got := s.LiveShells(); got != tt.want {
				t.Errorf("shotgun.LiveShells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shotgun_Shoot(t *testing.T) {
	type fields struct {
		shellsLeft uint8
		chamber    uint8
		dmg        uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		{
			name: "normal, live",
			fields: fields{
				shellsLeft: 1,
				chamber: 0b00000101,
				dmg: 1,
			},
			want: 1,
		},
		{
			name: "2x dmg, live",
			fields: fields{
				shellsLeft: 1,
				chamber: 0b00000101,
				dmg: 2,
			},
			want: 2,
		},
		{
			name: "normal, blank",
			fields: fields{
				shellsLeft: 1,
				chamber: 0,
				dmg: 1,
			},
			want: 0,
		},
		{
			name: "2x dmg, blank",
			fields: fields{
				shellsLeft: 1,
				chamber: 0,
				dmg: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shotgun{
				shellsLeft: tt.fields.shellsLeft,
				chamber:    tt.fields.chamber,
				dmg:        tt.fields.dmg,
			}
			if got := s.Shoot(); got != tt.want {
				t.Errorf("shotgun.Shoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
