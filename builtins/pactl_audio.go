package builtins

import (
	"barster/pkg"
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func pactlAudioUpdate() string {
	// get the sink volume in "55%" format
	getSinkVolumeCmd := "pactl get-sink-volume @DEFAULT_SINK@ | sed 's/.*\\/ \\+\\([0-9]\\+%\\) \\+\\/.*/\\1/' | head -n 1"
	sinkVolumePercent, err := exec.Command("bash", "-c", getSinkVolumeCmd).Output()
	if err != nil {
		return "ERR"
	}
	sinkVolumePercentStr := strings.TrimSpace(string(sinkVolumePercent))

	// get sink mute state
	getSinkMuteCmd := "pactl get-sink-mute @DEFAULT_SINK@ | sed 's/Mute: //'"
	sinkMuteStr, err := exec.Command("bash", "-c", getSinkMuteCmd).Output()
	if err != nil {
		return "ERR"
	}
	sinkMute := string(sinkMuteStr) == "yes"

	// get the source volume in "55%" format
	getSourceVolumeCmd := "pactl get-source-volume @DEFAULT_SOURCE@ | sed 's/.*\\/ \\+\\([0-9]\\+%\\) \\+\\/.*/\\1/' | head -n 1"
	sourceVolumePercent, err := exec.Command("bash", "-c", getSourceVolumeCmd).Output()
	if err != nil {
		return "ERR"
	}
	sourceVolumePercentStr := strings.TrimSpace(string(sourceVolumePercent))

	// get source mute state
	getSourceMuteCmd := "pactl get-source-mute @DEFAULT_SOURCE@ | sed 's/Mute: //'"
	sourceMuteStr, err := exec.Command("bash", "-c", getSourceMuteCmd).Output()
	if err != nil {
		return "ERR"
	}
	sourceMute := string(sourceMuteStr) == "yes"

	ret := ""
	if sinkMute {
		ret += "MUTE"
	} else {
		ret += fmt.Sprintf("%4s", sinkVolumePercentStr)
	}

	ret += "|"
	if sourceMute {
		ret += "MUT"
	} else {
		ret += fmt.Sprintf("%4s", sourceVolumePercentStr)
	}

	return ret
}

func startPactlSubscriber() chan struct{} {
	ch := make(chan struct{})
	go func() {
		cmd := exec.Command("pactl", "subscribe")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error starting pactl subscribe:", err)
			close(ch)
			return
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting pactl subscribe:", err)
			close(ch)
			return
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			// only lines that contain "change" are relevant
			if strings.Contains(line, "change") {
				ch <- struct{}{}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading pactl subscribe output:", err)
		}

		close(ch)
	}()

	return ch
}

func PactlAudioModule() pkg.Module {
	return pkg.Module{
		Name:     "PactlAudio",
		Interval: 24 * time.Hour,
		Update:   pactlAudioUpdate,
		Ticker:   startPactlSubscriber(),
	}
}
