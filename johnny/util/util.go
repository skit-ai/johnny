package util

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadColumnRow(pathToFile string) []string {

	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalf("Unable to read input file %v, %v\n", pathToFile, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	columnRow, err := csvReader.Read()
	if err != nil {
		log.Fatal("Unable to parse first row from CSV "+pathToFile, err)
	}

	return columnRow
}

func IdentifyAudioURLColumnPosition(columnRow []string) (bool, int) {

	audioUrlColumnNames := []string{"audio_url", "s3_audio_url"}

	for i, columnName := range columnRow {
		for _, columnConstant := range audioUrlColumnNames {
			if columnName == columnConstant {
				columnPos := i
				return true, columnPos
			}
		}
	}

	log.Fatalf("could not find audio urls in column names expected := %v\n"+
		"from given column names := %v\n", audioUrlColumnNames, columnRow)
	return false, 0

}

func ReadOnlyAudioURLs(pathToFile string, columnPos int) []string {

	var audioURL string
	audioURLs := []string{}

	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalf("Unable to read input file %v, %v\n", pathToFile, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1

	idx := 0
	for {
		// read just one record at a time
		record, err := csvReader.Read()
		idx += 1
		// fmt.Println(record)

		// to weed out rows which are less than what we expect in rows.
		if len(record) > columnPos {

			audioURL = record[columnPos]
			audioURLs = append(audioURLs, audioURL)

		}

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Unable to read %d row in csv file %v, %v\n", idx, pathToFile, err)
		}

	}

	// to skip column name, from first row.
	return audioURLs[1:]
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
