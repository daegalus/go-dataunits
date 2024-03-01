package units

import (
	"reflect"
	"testing"

	"github.com/daegalus/go-dataunits/units/memory"
	"github.com/daegalus/go-dataunits/units/storage"
)

func TestFromAbbreviation(t *testing.T) {
	type args struct {
		abbr string
	}
	type testValues[T Units] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	testsSU := []testValues[storage.StorageUnit]{
		{
			name: "storage KILOBYTE",
			args: args{
				abbr: "KB",
			},
			want:    storage.KILOBYTE,
			wantErr: false,
		},
		{
			name: "storage YOBIBYTE",
			args: args{
				abbr: "YiB",
			},
			want:    storage.YOBIBYTE,
			wantErr: false,
		},
		// Add more test cases for storage.StorageUnit if needed
	}

	testsMU := []testValues[memory.MemoryUnit]{
		{
			name: "memory KILOBYTE",
			args: args{
				abbr: "KB",
			},
			want:    memory.KILOBYTE,
			wantErr: false,
		},
		{
			name: "memory TERABYTE",
			args: args{
				abbr: "TB",
			},
			want:    memory.TERABYTE,
			wantErr: false,
		},
		// Add more test cases for memory.MemoryUnit if needed
	}

	for _, tt := range testsSU {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromAbbreviation[storage.StorageUnit](tt.args.abbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromAbbreviation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range testsMU {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromAbbreviation[memory.MemoryUnit](tt.args.abbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromAbbreviation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
