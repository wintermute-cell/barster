package main

import (
	"dwl_asyncbar/builtins"
	"dwl_asyncbar/pkg"
)

func main() {
	// Define the modules with their update intervals
	modules := []pkg.Module{
		builtins.NetTrafficModule(),
		builtins.BatteryModule(),
		builtins.DateTimeModule(),
	}

	// Separator between module outputs
	separator := " || "

	// Initialize the status bar
	statusBar := pkg.NewStatusBar(modules, separator)

	// Run the status bar
	statusBar.Start()

	// Ensure graceful termination
	// Optionally handle OS signals (e.g., SIGINT) if needed.
}
