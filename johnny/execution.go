package johnny

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"johnny/johnny/util"

	"github.com/schollz/progressbar/v3"
)

const (
	MAX_WORKERS = 30
	AUDIO_RATE  = "8k"
)

func Run() {

	start := time.Now()

	inputPathArg := flag.String("input", "input.csv", "csv which contains audio urls")
	outputWavDirArg := flag.String("output", "wav_audios", "directory where the wav audios need to be stored.")
	workersArg := flag.Int("workers", MAX_WORKERS, "maximum goroutines in the pool")
	audioRateArg := flag.String("rate", AUDIO_RATE, "audio sample rate / frequency of output audios.")
	flag.Parse()

	csvPath := *inputPathArg
	outputWavDir := *outputWavDirArg
	workers := *workersArg
	rate := *audioRateArg

	util.CreateDir(outputWavDir)

	columnRow := util.ReadColumnRow(csvPath)
	_, columnPos := util.IdentifyAudioURLColumnPosition(columnRow)

	audioURLs := make(chan string, util.Max(workers*2, 100))
	jobs := make(chan Job, workers)

	var wg sync.WaitGroup
	bar := progressbar.Default(int64(-1), "downloading & converting audios ...")

	noOfAudioURLs := 0
	wg.Add(1)

	go func() {

		defer wg.Done()
		defer close(jobs)
		defer bar.Finish()

		for audioURL := range audioURLs {
			jobs <- Job{AudioURL: audioURL, WavAudioDirPath: outputWavDir, AudioRate: rate}

			noOfAudioURLs += 1
			bar.Add(1)
		}

	}()

	go util.ReadOnlyAudioURLs(csvPath, columnPos, audioURLs)

	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func(id int) {
			worker(id, jobs)
			wg.Done()
		}(i)
	}

	wg.Wait()

	timeLapsed := time.Since(start)
	// fmt.Printf("->> johnny finished downloading & converting %d audios to %vHz under %v, stored them in %v.\n", noOfAudioURLs, rate, timeLapsed, outputWavDir)

	fmt.Printf("->> johnny took %v for %d audios. they are stored under the directory: %v.\n", timeLapsed, noOfAudioURLs, outputWavDir)

}
