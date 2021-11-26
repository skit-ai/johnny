package util

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(pathToFile string) [][]string {

	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalf("Unable to read input file %v, %v\n", pathToFile, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+pathToFile, err)
	}

	return records
}

func identifyColumnPosition(records [][]string) (bool, int) {

	audioUrlColumnNames := []string{"audio_url", "s3_audio_url"}

	for i, columnName := range records[0] {
		for _, columnConstant := range audioUrlColumnNames {
			if columnName == columnConstant {
				columnPos := i
				return true, columnPos
			}
		}
	}

	log.Fatalf("could not find audio urls := %v\n", audioUrlColumnNames)
	return false, 0

}

func ExtractAudioURLs(records [][]string) []string {

	_, columnPos := identifyColumnPosition(records)

	audios := []string{}

	// starts at 1 to skip column names.
	// assumes second column is audio url, and takes it.
	for i := 1; i < len(records); i++ {
		audios = append(audios, records[i][columnPos])
	}

	return audios

}

func CreateDir(directoryPath string) {

	err := os.MkdirAll(directoryPath, os.ModePerm)

	if err != nil {
		log.Fatalf("Not able to create directory %v, %v", directoryPath, err)
	}

}

func DeleteTmpFile(tmpFileName string) {

	err := os.Remove(tmpFileName)
	if err != nil {
		log.Println(err)
	}

}
