package main

import (
	"barster/builtins"
	"barster/pkg"

	"github.com/godbus/dbus/v5"
)

func getTrayIconCount() int {
	conn, err := dbus.SessionBus()
	if err != nil {
		return 0
	}
	defer conn.Close()

	variant, err := conn.Object("org.kde.StatusNotifierWatcher", "/StatusNotifierWatcher").
		GetProperty("org.kde.StatusNotifierWatcher.RegisteredStatusNotifierItems")
	if err != nil {
		return 0
	}

	// Ensure the value is an array of strings.
	registeredItems, ok := variant.Value().([]string)
	if !ok {
		return 0
	}

	return len(registeredItems)
}

func main() {
	// define your list of modules here, from right to left
	modules := []pkg.Module{
		builtins.PactlAudioModule(),
		builtins.NetTrafficModule(),
		builtins.BatteryModule(),
		builtins.DateTimeModule(),
		builtins.DynamicSpacerModule(func() int {
			return getTrayIconCount() * 2
		}),
	}

	// choose a separator character that is placed between each module
	separator := " ║ "

	// here, the bar is created and started, you can ignore this as a user
	statusBar := pkg.NewStatusBar(modules, separator)
	statusBar.Start()
}
