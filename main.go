package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	countBytesFlag := flag.Bool("c", false, "a boolean flag for counting the number of bytes")

	flag.Parse()

	filePath := flag.Arg(0)

	fileInfo := getFileInfo(filePath)

	if *countBytesFlag {
		log.Printf("%d %s", fileInfo.Size(), filePath)
	}
}

func getFileInfo(filePath string) os.FileInfo {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo
}
