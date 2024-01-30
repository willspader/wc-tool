// TODO: Tests for large files (50GB)
// TODO: If failed (take too long or out of memory) for large files, then take another path to not stored all the file in memory.
// If file smaller than some threshhold, then load to the memory, else, iterate over the file without loading all the content.

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Flags struct {
	bytesCounterFlag bool
	linesCounterFlag bool
}

type Counters struct {
	bytesCounter int
	linesCounter int
}

func main() {
	flags := parseFlags()

	filePath := flag.Arg(0)

	file := getFile(filePath)

	counters := resolve(file, flags)

	output(filePath, flags, counters)
}

func getFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func countLines(file []byte) int {
	total := 0
	for i := range file {
		if file[i] == 10 {
			total++
		}
	}
	return total
}

func countBytes(file []byte) int {
	return len(file)
}

func countWords(file []byte) int {
	return 0
}

func output(filePath string, flags Flags, counters Counters) {
	output := ""

	if flags.bytesCounterFlag {
		output = output + fmt.Sprint(counters.bytesCounter)
	}

	if flags.linesCounterFlag {
		output = output + " " + fmt.Sprint(counters.linesCounter)
	}

	output = output + " " + filePath

	fmt.Println(strings.TrimPrefix(output, " "))
}

func parseFlags() Flags {
	bytesCounterFlag := flag.Bool("c", false, "a boolean flag for counting the number of bytes")
	linesCounterFlag := flag.Bool("l", false, "a boolean flag for counting the number of lines")

	flag.Parse()

	// if none provided, then show all
	if !*bytesCounterFlag && !*linesCounterFlag {
		*bytesCounterFlag = true
		*linesCounterFlag = true
	}

	return Flags{bytesCounterFlag: *bytesCounterFlag, linesCounterFlag: *linesCounterFlag}
}

func resolve(file []byte, flags Flags) Counters {
	outputs := newOutputs()
	if flags.bytesCounterFlag {
		outputs.bytesCounter = countBytes(file)
	}

	if flags.linesCounterFlag {
		outputs.linesCounter = countLines(file)
	}

	return outputs
}

func newOutputs() Counters {
	return Counters{}
}
