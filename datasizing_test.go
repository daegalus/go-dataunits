package dataunits

import (
	"reflect"
	"testing"

	"github.com/daegalus/go-dataunits/units/memory"
	"github.com/daegalus/go-dataunits/units/storage"
)

func TestParseSizeStorageUnit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    DataSize[storage.StorageUnit]
		wantErr bool
	}{
		{
			name:    "Valid input - Kilobyte",
			args:    args{input: "1KB"},
			want:    DataSize[storage.StorageUnit]{Size: 1, Unit: storage.KILOBYTE, Bytes: 1000},
			wantErr: false,
		},
		{
			name:    "Valid input - Mebibyte",
			args:    args{input: "2MiB"},
			want:    DataSize[storage.StorageUnit]{Size: 2, Unit: storage.MEBIBYTE, Bytes: 2 * 1024 * 1024},
			wantErr: false,
		},
		{
			name:    "Valid input - Gigabyte",
			args:    args{input: "3GB"},
			want:    DataSize[storage.StorageUnit]{Size: 3, Unit: storage.GIGABYTE, Bytes: 3 * 1000 * 1000 * 1000},
			wantErr: false,
		},
		{
			name:    "Invalid input - Missing unit",
			args:    args{input: "10"},
			want:    DataSize[storage.StorageUnit]{},
			wantErr: true,
		},
		{
			name:    "Invalid input - Unknown unit",
			args:    args{input: "5SB"},
			want:    DataSize[storage.StorageUnit]{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSize[storage.StorageUnit](tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSize() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestParseSizeMemoryUnit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    DataSize[memory.MemoryUnit]
		wantErr bool
	}{
		{
			name:    "Valid input - Kilobyte",
			args:    args{input: "1KB"},
			want:    DataSize[memory.MemoryUnit]{Size: 1, Unit: memory.KILOBYTE, Bytes: 1024},
			wantErr: false,
		},
		{
			name:    "Invalid input - Mebibyte",
			args:    args{input: "2MiB"},
			want:    DataSize[memory.MemoryUnit]{},
			wantErr: true,
		},
		{
			name:    "Valid input - Gigabyte",
			args:    args{input: "3GB"},
			want:    DataSize[memory.MemoryUnit]{Size: 3, Unit: memory.GIGABYTE, Bytes: 3 * 1024 * 1024 * 1024},
			wantErr: false,
		},
		{
			name:    "Invalid input - Missing unit",
			args:    args{input: "10"},
			want:    DataSize[memory.MemoryUnit]{},
			wantErr: true,
		},
		{
			name:    "Invalid input - Unknown unit",
			args:    args{input: "5PB"},
			want:    DataSize[memory.MemoryUnit]{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSize[memory.MemoryUnit](tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSize() = %v, want %v", got, tt.want)
			}
		})
	}

}
