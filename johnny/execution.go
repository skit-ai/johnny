package johnny

import (
	"flag"
	"fmt"
	"sync"
	
	"johnny/johnny/util"

	"github.com/schollz/progressbar/v3"
)

const (
	MAX_WORKERS = 30
)



func Run() {

	inputPathArg := flag.String("input", "input.csv", "csv which contains audio urls")
	outputWavDirArg := flag.String("output", "wav_audios", "directory where the wav audios need to be stored.")
	workersArg := flag.Int("workers", MAX_WORKERS, "maximum goroutines in the pool")
	flag.Parse()

	csvPath := *inputPathArg
	outputWavDir := * outputWavDirArg
	workers := *workersArg

	util.CreateDir(outputWavDir)
	
	records := util.ReadCsvFile(csvPath)
	audioURLs := util.ExtractAudioURLs(records)

	jobs := make(chan Job, workers)
	var wg sync.WaitGroup
	bar := progressbar.Default(int64(len(audioURLs)))


	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func(id int) {
			worker(id, jobs)
			wg.Done()
		}(i)
	}

	for _, audioURL := range audioURLs {
		jobs <- Job{AudioURL: audioURL, WavAudioDirPath: outputWavDir}	
		bar.Add(1)
	}
	close(jobs)

	wg.Wait()
	fmt.Printf("finished processing %d audios and stored them in %v\n", len(audioURLs), outputWavDir)

}