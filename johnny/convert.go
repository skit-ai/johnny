package johnny

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func isRaw(inputAudioString string) bool {

	return strings.HasSuffix(inputAudioString, ".raw")

}

func getFfmpegCommandArgs(inputAudioPath string, outputAudioPath string, outputRate string, isRawFile bool) []string {

	var ffmpegArgs []string

	if isRawFile {
		ffmpegArgs = []string{
			"-f", "s16le",
			"-ar", "8k",
			"-ac", "1",
			"-i", inputAudioPath,
			"-acodec", "pcm_s16le",
			"-ac", "1",
			"-ar", outputRate,
			outputAudioPath,
			"-hide_banner",
		}
	} else {
		ffmpegArgs = []string{
			"-i", inputAudioPath,
			"-acodec", "pcm_s16le",
			"-ac", "1",
			"-ar", outputRate,
			outputAudioPath,
			"-hide_banner",
		}
	}

	return ffmpegArgs

}

func convertAudioToSpecificKiloHz(inputAudioPath string, outputAudioPath string, rate string, isRawFile bool) {

	var out bytes.Buffer
	var stderr bytes.Buffer

	name := "ffmpeg"
	args := getFfmpegCommandArgs(inputAudioPath, outputAudioPath, rate, isRawFile)
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
