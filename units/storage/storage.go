package storage

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	BINARY  = 1024
	DECIMAL = 1000
	BYTE    = 1

	KILOBYTE  StorageUnit = DECIMAL * BYTE
	MEGABYTE              = DECIMAL * KILOBYTE
	GIGABYTE              = DECIMAL * MEGABYTE
	TERABYTE              = DECIMAL * GIGABYTE
	PETABYTE              = DECIMAL * TERABYTE
	EXABYTE               = DECIMAL * PETABYTE
	ZETTABYTE             = DECIMAL * EXABYTE
	YOTTABYTE             = DECIMAL * ZETTABYTE

	KIBIBYTE StorageUnit = BINARY * BYTE
	MEBIBYTE             = BINARY * KIBIBYTE
	GIBIBYTE             = BINARY * MEBIBYTE
	TEBIBYTE             = BINARY * GIBIBYTE
	PEBIBYTE             = BINARY * TEBIBYTE
	EXBIBYTE             = BINARY * PEBIBYTE
	ZEBIBYTE             = BINARY * EXBIBYTE
	YOBIBYTE             = BINARY * ZEBIBYTE
)

type StorageUnit float64

func IsStorageUnit[T any]() bool {
	return reflect.TypeFor[T]().Name() == reflect.TypeFor[StorageUnit]().Name()
}

func FromAbbreviation(abbr string) (StorageUnit, error) {
	abbr = strings.TrimSuffix(strings.ToLower(abbr), "b")
	var unit StorageUnit
	var err error

	switch abbr {
	case "":
		unit = BYTE
	case "k":
		unit = KILOBYTE
	case "m":
		unit = MEGABYTE
	case "g":
		unit = GIGABYTE
	case "t":
		unit = TERABYTE
	case "p":
		unit = PETABYTE
	case "e":
		unit = EXABYTE
	case "z":
		unit = ZETTABYTE
	case "y":
		unit = YOTTABYTE
	case "ki":
		unit = KIBIBYTE
	case "mi":
		unit = MEBIBYTE
	case "gi":
		unit = GIBIBYTE
	case "ti":
		unit = TEBIBYTE
	case "pi":
		unit = PEBIBYTE
	case "ei":
		unit = EXBIBYTE
	case "zi":
		unit = ZEBIBYTE
	case "yi":
		unit = YOBIBYTE
	default:
		err = fmt.Errorf("unknown abbreviation: %s", abbr)
	}

	return unit, err
}
