package builtins

import (
	"dwl_asyncbar/pkg"
	"fmt"
	"os"
	"time"

	"github.com/distatus/battery"
)

var statusStrings = map[int8]string{
	-1: "ERR",
	0:  "ERR",
	1:  "E",
	2:  "F",
	3:  "+",
	4:  "-",
	5:  "I",
}

func batteryStatus() string {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get battery info: %v\n", err)
		return "Error"
	}
	ret := ""
	for i, b := range batteries {
		if i > 0 {
			ret += ", "
		}
		ret += fmt.Sprintf("%s%.0f%%", statusStrings[int8(b.State.Raw)], b.Current/b.Full*100)
	}

	return ret
}

// BatteryModule returns a prebuilt Battery module.
//
// Status legend:
//
// ERR: Error
//
// E: Empty
//
// F: Full
//
// +: Charging
//
// -: Discharging
//
// I: Idle
func BatteryModule() pkg.Module {
	return pkg.Module{
		Name:     "Battery",
		Interval: 20 * time.Second,
		Update:   batteryStatus,
	}
}
