package builtins

import (
	"barster/pkg"
	"math/rand"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Scuffed UID, but good enough, lets us avoid uuid dependency
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// SpacerModule creates a module that outputs a fixed amount of spaces.
func SpacerModule(width int) pkg.Module {
	return pkg.Module{
		Name:     "Spacer" + randSeq(8),
		Interval: 24 * time.Hour, // This doesnt really have to update
		Update: func() string {
			return strings.Repeat(" ", width)
		},
	}
}

// DynamicSpacerModule creates a module that outputs a dynamic amount of
// spaces. The spaceFunc function is called every second to determine the
// amount of spaces to output.
func DynamicSpacerModule(spaceFunc func() int) pkg.Module {
	return pkg.Module{
		Name:     "DynamicSpacer" + randSeq(8),
		Interval: 1 * time.Second,
		Update: func() string {
			return strings.Repeat(" ", spaceFunc())
		},
	}
}
