package main

import (
	"log"
	"os"
)

func HandleSplit(file *os.File, fileCount int, delim string, clean bool) int {
	var records [][]string

	var err error
	if clean {
		records, err = ReadAndCleanCSVFile(file, rune(delim[0]))
	} else {
		records, err = ReadCSVFile(file, rune(delim[0]))
	}

	if err != nil {
		log.Fatalf("error opening file %s: %v\n", file.Name(), err)
	}

	SetHeaders(records[0])
	records = records[1:]
	SplitFiles(records, file.Name(), fileCount, rune(delim[0]))

	return fileCount
}
