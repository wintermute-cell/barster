package builtins

import (
	"barster/pkg"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func mtrakkerUpdate() string {
	url := "http://localhost:8080/current-track"

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "ERR"
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "ERR"
	}

	// Parse the JSON response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "ERR"
	}

	// Get the value of the "currentTrack" field
	currentTrack, ok := result["currentTrack"].(string)
	if !ok {
		return "ERR"
	}

	return "♫ " + currentTrack
}

// MtrakkerModule returns a module that calls mtrakker to get the currently
// playing youtube song.
func MtrakkerModule() pkg.Module {
	return pkg.Module{
		Name:     "Mtrakker",
		Interval: 1 * time.Second,
		Update:   mtrakkerUpdate,
	}
}
