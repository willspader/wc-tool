// TODO: Tests for large files (50GB)
// TODO: If failed (take too long or out of memory) for large files, then take another path to not stored all the file in memory.
// If file smaller than some threshhold, then load to the memory, else, iterate over the file without loading all the content.

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type FlagCounter struct {
	active    bool
	counter   int
	calculate func([]byte) int
	print     func(string)
}

func main() {
	flagsCounters := parseFlags()

	filePath := flag.Arg(0)

	file := getFile(filePath)

	resolveCalculations(file, flagsCounters)

	resolveOutput(filePath, flagsCounters)
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
	return len(strings.Fields(string(file)))
}

// If the current locale does not support multibyte characters this will match the -c option.
func countCharacters(file []byte) int {
	return utf8.RuneCount(file)
}

func resolveCalculations(file []byte, flagsCounters []FlagCounter) {
	for i := 0; i < len(flagsCounters); i++ {
		if flagsCounters[i].active {
			flagsCounters[i].counter = flagsCounters[i].calculate(file)
		}
	}
}

func resolveOutput(filePath string, flagsCounters []FlagCounter) {
	output := ""

	for _, flagCounter := range flagsCounters {
		if flagCounter.active {
			output = output + " " + fmt.Sprint(flagCounter.counter)
		}
	}

	output = output + " " + filePath

	fmt.Println(strings.TrimPrefix(output, " "))
}

func parseFlags() []FlagCounter {
	bytesCounterFlag := flag.Bool("c", false, "a boolean flag for counting the number of bytes")
	linesCounterFlag := flag.Bool("l", false, "a boolean flag for counting the number of lines")
	wordsCounterFlag := flag.Bool("w", false, "a boolean flag for counting the number of words")
	charactersCounterFlag := flag.Bool("m", false, "a boolean flag for counting the number of characters")

	flag.Parse()

	// if none provided, then show -c -l -w
	if !*bytesCounterFlag && !*linesCounterFlag && !*wordsCounterFlag && !*charactersCounterFlag {
		*bytesCounterFlag = true
		*linesCounterFlag = true
		*wordsCounterFlag = true
	}

	return []FlagCounter{
		{active: *linesCounterFlag, calculate: countLines},
		{active: *wordsCounterFlag, calculate: countWords},
		{active: *bytesCounterFlag, calculate: countBytes},
		{active: *charactersCounterFlag, calculate: countCharacters}}
}
