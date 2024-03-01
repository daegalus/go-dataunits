package dataunits

import (
	"fmt"
	"log/slog"
	"regexp"
	"strconv"

	"github.com/daegalus/go-dataunits/units"
	"github.com/daegalus/go-dataunits/units/memory"
	"github.com/daegalus/go-dataunits/units/storage"
)

type DataSize[T units.Units] struct {
	Size  float64
	Unit  T
	Bytes float64
}

func ParseSize[T units.Units](input string) (DataSize[T], error) {
	regex, err := regexp.Compile("([0-9]+[.]?[0-9]*)([kbmgtpezyKBMGTPEZYiI].{1,2})")
	if err != nil {
		return DataSize[T]{}, err
	}

	matches := regex.FindStringSubmatch(input)
	slog.Info("matches", slog.Any("matches", matches))
	if len(matches) < 3 {
		return DataSize[T]{}, fmt.Errorf("provided input is not a valid format")
	}

	sizeStr := matches[1]
	unitStr := matches[2]

	size, err := strconv.ParseFloat(sizeStr, 64)
	if err != nil {
		return DataSize[T]{}, err
	}

	unit, err := units.FromAbbreviation[T](unitStr)
	if err != nil {
		return DataSize[T]{}, err
	}

	return DataSize[T]{
		Size:  size,
		Unit:  unit,
		Bytes: size * float64(unit),
	}, nil
}

func (ds DataSize[T]) Kilobyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.KILOBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(memory.KILOBYTE)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Megabyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.MEGABYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(memory.MEGABYTE)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Gigabyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.GIGABYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(memory.GIGABYTE)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Terabyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.TERABYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(memory.TERABYTE)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Petabyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.PETABYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Exabyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.EXABYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Zettabyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.ZETTABYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Yottabyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.YOTTABYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Kibibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.KIBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Mebibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.MEBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Gibibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.GIBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Tebibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.TEBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Pebibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.PEBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Exbibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.EXBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Zebibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.ZEBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}

func (ds DataSize[T]) Yobibyte() float64 {
	var unit T
	if storage.IsStorageUnit[T]() {
		unit = T(storage.YOBIBYTE)
	} else if memory.IsMemoryUnit[T]() {
		unit = T(1)
	}
	return ds.Bytes / float64(unit)
}
