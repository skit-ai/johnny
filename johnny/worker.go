package johnny

import (
	"log"
	"path/filepath"
	"strings"

	"johnny/johnny/util"
)

func worker(id int, input chan Job) {

	for job := range input {

		url := job.AudioURL

		pathSplit := strings.Split(url, "/")
		nameOfAudioFile := pathSplit[len(pathSplit)-1]

		// task 1
		tmpFileName, err := downloadAndSaveFile(url)
		if err != nil {
			log.Println(err)
			return
		}

		fileNameSplit := strings.Split(nameOfAudioFile, ".")
		wav8khzConvertedFilename := strings.Join(fileNameSplit[:len(fileNameSplit)-1], ".") + ".wav"
		convertedAudioFilePath := filepath.Join(job.WavAudioDirPath, wav8khzConvertedFilename)

		// task 2
		convertAudioTo8hz(tmpFileName, convertedAudioFilePath)

		util.DeleteTmpFile(tmpFileName)

	}

}
