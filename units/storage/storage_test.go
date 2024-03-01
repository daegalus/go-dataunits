package storage

import (
	"testing"

	"github.com/daegalus/go-dataunits/units/memory"
	"github.com/stretchr/testify/assert"
)

func TestFromAbbreviation(t *testing.T) {
	type args struct {
		abbr string
	}
	tests := []struct {
		name    string
		args    args
		want    StorageUnit
		wantErr bool
	}{
		{
			name: "Test abbreviation KB",
			args: args{abbr: "KB"},
			want: KILOBYTE,
		},
		{
			name: "Test abbreviation MB",
			args: args{abbr: "MB"},
			want: MEGABYTE,
		},
		{
			name: "Test abbreviation GB",
			args: args{abbr: "GB"},
			want: GIGABYTE,
		},
		{
			name: "Test abbreviation TB",
			args: args{abbr: "TB"},
			want: TERABYTE,
		},
		{
			name: "Test abbreviation PB",
			args: args{abbr: "PB"},
			want: PETABYTE,
		},
		{
			name: "Test abbreviation EB",
			args: args{abbr: "EB"},
			want: EXABYTE,
		},
		{
			name: "Test abbreviation ZB",
			args: args{abbr: "ZB"},
			want: ZETTABYTE,
		},
		{
			name: "Test abbreviation YB",
			args: args{abbr: "YB"},
			want: YOTTABYTE,
		},
		{
			name: "Test abbreviation KiB",
			args: args{abbr: "KiB"},
			want: KIBIBYTE,
		},
		{
			name: "Test abbreviation MiB",
			args: args{abbr: "MiB"},
			want: MEBIBYTE,
		},
		{
			name: "Test abbreviation GiB",
			args: args{abbr: "GiB"},
			want: GIBIBYTE,
		},
		{
			name: "Test abbreviation TiB",
			args: args{abbr: "TiB"},
			want: TEBIBYTE,
		},
		{
			name: "Test abbreviation PiB",
			args: args{abbr: "PiB"},
			want: PEBIBYTE,
		},
		{
			name: "Test abbreviation EiB",
			args: args{abbr: "EiB"},
			want: EXBIBYTE,
		},
		{
			name: "Test abbreviation ZiB",
			args: args{abbr: "ZiB"},
			want: ZEBIBYTE,
		},
		{
			name: "Test abbreviation YiB",
			args: args{abbr: "YiB"},
			want: YOBIBYTE,
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

func TestIsStorageUnit(t *testing.T) {
	t.Run("MemoryUnit is not a StorageUnit", func(t *testing.T) {
		assert.False(t, IsStorageUnit[memory.MemoryUnit]())
	})

	t.Run("StorageUnit is a StorageUnit", func(t *testing.T) {
		assert.True(t, IsStorageUnit[StorageUnit]())
	})
}
