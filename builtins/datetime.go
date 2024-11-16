package builtins

import (
	"dwl_asyncbar/pkg"
	"time"
)

func dateTimeWithFormat(format string) func() string {
	return func() string {
		return time.Now().Format(format)
	}
}

// DateTimeModule returns a prebuilt DateTime module.
//
// Displays:  Mon 02. Jan 15:04:05 2006
func DateTimeModule() pkg.Module {
	return pkg.Module{
		Name:     "DateTime",
		Interval: 1 * time.Second,
		Update:   dateTimeWithFormat("Mon 02. Jan 15:04:05 2006"),
	}
}

// DateTimeModuleFormat returns a DateTime module with a custom format.
// Format follows the same rules as time.Format.
func DateTimeModuleFormat(format string) pkg.Module {
	return pkg.Module{
		Name:     "DateTime",
		Interval: 1 * time.Second,
		Update:   dateTimeWithFormat(format),
	}
}
