package johnny

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func convertAudioTo8hz(inputAudioPath string, outputAudioPath string) {

	var out bytes.Buffer
	var stderr bytes.Buffer

	name := "ffmpeg"
	args := []string{"-i", inputAudioPath, "-acodec", "pcm_s16le", "-ac", "1", "-ar", "8k", outputAudioPath}
	cmd := exec.Command(name, args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if !strings.Contains(stderr.String(), "Overwrite") {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		}
		return
	}

}
