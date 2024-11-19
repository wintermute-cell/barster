package main

import (
	"barster/builtins"
	"barster/pkg"
)

func main() {
	// define your list of modules here, from right to left
	modules := []pkg.Module{
		builtins.NetTrafficModule(),
		builtins.BatteryModule(),
		builtins.DateTimeModule(),
	}

	// choose a separator character that is placed between each module
	separator := " ║ "

	// here, the bar is created and started, you can ignore this as a user
	statusBar := pkg.NewStatusBar(modules, separator)
	statusBar.Start()
}
