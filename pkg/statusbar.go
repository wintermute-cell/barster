package pkg

import (
	"fmt"
	"sync"
	"time"
)

// StatusBar handles updating and combining module outputs.
type StatusBar struct {
	Modules   []Module
	Updates   chan string
	Separator string
	outputs   map[string]string
	mu        sync.Mutex
}

// NewStatusBar initializes a new StatusBar.
func NewStatusBar(modules []Module, separator string) *StatusBar {
	return &StatusBar{
		Modules:   modules,
		Updates:   make(chan string),
		Separator: separator,
		outputs:   make(map[string]string),
	}
}

// runModule starts a module's periodic updates and sends results to the StatusBar.
func (sb *StatusBar) runModule(module Module, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(module.Interval)
	defer ticker.Stop()

	if module.Ticker == nil {
		module.Ticker = make(chan struct{})
	}

	runModuleInternal := func() {
		output := module.Update()
		sb.mu.Lock()
		sb.outputs[module.Name] = output
		sb.mu.Unlock()
		sb.refresh()
	}

	runModuleInternal() // Run once immediately

	for {
		select {
		case <-ticker.C:
			runModuleInternal()
		case <-module.Ticker:
			runModuleInternal()
		}
	}
}

// refresh concatenates all module outputs and sends the result to stdout.
func (sb *StatusBar) refresh() {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	var combinedOutput string
	for i, module := range sb.Modules {
		combinedOutput += sb.outputs[module.Name]
		if i < len(sb.Modules)-1 {
			combinedOutput += sb.Separator
		}
	}
	// Print the combined output to stdout.
	fmt.Println(combinedOutput)
}

// Start initializes and runs all modules.
func (sb *StatusBar) Start() {
	var wg sync.WaitGroup
	for _, module := range sb.Modules {
		wg.Add(1)
		go sb.runModule(module, &wg)
	}
	wg.Wait() // Wait indefinitely for all modules (they won't stop unless modified).
}
