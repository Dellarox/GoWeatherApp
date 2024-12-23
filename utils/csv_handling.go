package utils

import (
	"encoding/csv"
	"os"
)

func CreateCSVWriter(filename string) (*csv.Writer, *os.File) {
	var header = []string{"location", "temperature", "description", "humidity", "wind_speed", "coordinates", "time"}
	var f *os.File
	var err error

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		CheckError(err)

		writer := csv.NewWriter(f)
		err = writer.Write(header)
		CheckError(err)
		writer.Flush()
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		CheckError(err)
	}

	writer := csv.NewWriter(f)
	return writer, f
}

func WriteCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	CheckError(err)
}
