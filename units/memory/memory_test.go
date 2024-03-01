package memory

import (
	"testing"

	"github.com/daegalus/go-dataunits/units/storage"
	"github.com/stretchr/testify/assert"
)

func TestFromAbbreviation(t *testing.T) {
	type args struct {
		abbr string
	}
	tests := []struct {
		name    string
		args    args
		want    MemoryUnit
		wantErr bool
	}{
		{
			name:    "Test Kilobyte",
			args:    args{abbr: "KB"},
			want:    KILOBYTE,
			wantErr: false,
		},
		{
			name:    "Test Megabyte",
			args:    args{abbr: "MB"},
			want:    MEGABYTE,
			wantErr: false,
		},
		{
			name:    "Test Gigabyte",
			args:    args{abbr: "GB"},
			want:    GIGABYTE,
			wantErr: false,
		},
		{
			name:    "Test Terabyte",
			args:    args{abbr: "TB"},
			want:    TERABYTE,
			wantErr: false,
		},
		{
			name:    "Test SupaByte (expect failure)",
			args:    args{abbr: "SB"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromAbbreviation(tt.args.abbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromAbbreviation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMemoryUnit(t *testing.T) {
	t.Run("StorageUnit is not a MemoryUnit", func(t *testing.T) {
		assert.False(t, IsMemoryUnit[storage.StorageUnit]())
	})

	t.Run("MemoryUnit is a MemoryUnit", func(t *testing.T) {
		assert.True(t, IsMemoryUnit[MemoryUnit]())
	})
}
