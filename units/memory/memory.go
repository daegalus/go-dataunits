package memory

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	BINARY = 1024
	BYTE   = 1

	KILOBYTE MemoryUnit = BINARY * BYTE
	MEGABYTE            = BINARY * KILOBYTE
	GIGABYTE            = BINARY * MEGABYTE
	TERABYTE            = BINARY * GIGABYTE
)

type MemoryUnit float64

func IsMemoryUnit[T any]() bool {
	return reflect.TypeFor[T]().Name() == reflect.TypeFor[MemoryUnit]().Name()
}

func FromAbbreviation(abbr string) (MemoryUnit, error) {
	abbr = strings.TrimSuffix(strings.ToLower(abbr), "b")
	switch abbr {
	case "":
		return BYTE, nil
	case "k":
		return KILOBYTE, nil
	case "m":
		return MEGABYTE, nil
	case "g":
		return GIGABYTE, nil
	case "t":
		return TERABYTE, nil
	default:
		return 0, fmt.Errorf("unknown abbreviation: %s", abbr)
	}
}
