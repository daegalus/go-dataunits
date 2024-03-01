package units

import (
	"fmt"
	"reflect"

	"github.com/daegalus/go-dataunits/units/memory"
	"github.com/daegalus/go-dataunits/units/storage"
)

type Units interface {
	storage.StorageUnit | memory.MemoryUnit
}

func FromAbbreviation[T Units](abbr string) (T, error) {
	if reflect.TypeFor[T]().Name() == reflect.TypeFor[storage.StorageUnit]().Name() {
		unit, err := storage.FromAbbreviation(abbr)
		if err != nil {
			return 0, err
		}
		return T(unit), nil
	} else if reflect.TypeFor[T]().Name() == reflect.TypeFor[memory.MemoryUnit]().Name() {
		unit, err := memory.FromAbbreviation(abbr)
		if err != nil {
			return 0, err
		}
		return T(unit), nil
	} else {
		return 0, fmt.Errorf("invalid unit type")
	}
}
