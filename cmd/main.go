package main

import (
	"log/slog"

	"github.com/daegalus/go-dataunits"
	"github.com/daegalus/go-dataunits/units/storage"
)

func main() {
	test := "4.08Mi"
	ds, _ := dataunits.ParseSize[storage.StorageUnit](test)
	slog.Info("The object", slog.Any("datasize", ds))

	test2 := "4.665432374687YiB"
	ds2, _ := dataunits.ParseSize[storage.StorageUnit](test2)
	slog.Info(
		"the sizes",
		slog.Float64("KB", ds2.Kilobyte()),
		slog.Float64("MB", ds2.Megabyte()),
		slog.Float64("GB", ds2.Gigabyte()),
		slog.Float64("TB", ds2.Terabyte()),
		slog.Float64("PB", ds2.Petabyte()),
		slog.Float64("EB", ds2.Exabyte()),
		slog.Float64("ZB", ds2.Zettabyte()),
		slog.Float64("YB", ds2.Yottabyte()),
		slog.Float64("KiB", ds2.Kibibyte()),
		slog.Float64("MiB", ds2.Mebibyte()),
		slog.Float64("GiB", ds2.Gibibyte()),
		slog.Float64("TiB", ds2.Tebibyte()),
		slog.Float64("PiB", ds2.Pebibyte()),
		slog.Float64("EiB", ds2.Exbibyte()),
		slog.Float64("ZiB", ds2.Zebibyte()),
		slog.Float64("YiB", ds2.Yobibyte()),
	)

	test3 := "4.665432374687yb"
	ds3, _ := dataunits.ParseSize[storage.StorageUnit](test3)
	slog.Info(
		"the sizes",
		slog.Float64("KB", ds3.Kilobyte()),
		slog.Float64("MB", ds3.Megabyte()),
		slog.Float64("GB", ds3.Gigabyte()),
		slog.Float64("TB", ds3.Terabyte()),
		slog.Float64("PB", ds3.Petabyte()),
		slog.Float64("EB", ds3.Exabyte()),
		slog.Float64("ZB", ds3.Zettabyte()),
		slog.Float64("YB", ds3.Yottabyte()),
		slog.Float64("KiB", ds3.Kibibyte()),
		slog.Float64("MiB", ds3.Mebibyte()),
		slog.Float64("GiB", ds3.Gibibyte()),
		slog.Float64("TiB", ds3.Tebibyte()),
		slog.Float64("PiB", ds3.Pebibyte()),
		slog.Float64("EiB", ds3.Exbibyte()),
		slog.Float64("ZiB", ds3.Zebibyte()),
		slog.Float64("YiB", ds3.Yobibyte()),
	)
}
