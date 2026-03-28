package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hasflag(flags []string, flag string) bool {
	for _, f := range flags {
		if f == flag {
			return true
		}
	}
	return false
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type searchFlags struct {
	hasLineNumbers    bool
	hasFileNames      bool
	isCaseInsensitive bool
	isInverted        bool
	isEntireLine      bool
	multipleFiles     bool
}

func findPattern(pattern string, lines []string, filename string, flags searchFlags) []string {
	var matches []string

	if flags.isCaseInsensitive {
		pattern = strings.ToLower(pattern)
	}
	for lineNumber, line := range lines {
		line = strings.TrimRight(line, " ")

		isMatch := false

		compareLine := line
		if flags.isCaseInsensitive {
			compareLine = strings.ToLower(compareLine)
		}

		if flags.isEntireLine {
			isMatch = compareLine == pattern
		} else {
			isMatch = strings.Contains(compareLine, pattern)
		}

		if flags.isInverted {
			isMatch = !isMatch
		}

		if isMatch {
			matchedLine := line

			if flags.hasLineNumbers {
				matchedLine = fmt.Sprintf("%d:%s", lineNumber+1, matchedLine)
			}

			if flags.multipleFiles {
				matchedLine = filename + ":" + matchedLine
			}
			matches = append(matches, matchedLine)
		}
	}

	return matches
}

func Search(pattern string, flags, files []string) []string {

	flagsStruct := searchFlags{
		hasLineNumbers:    hasflag(flags, "-n"),
		hasFileNames:      hasflag(flags, "-l"),
		isCaseInsensitive: hasflag(flags, "-i"),
		isInverted:        hasflag(flags, "-v"),
		isEntireLine:      hasflag(flags, "-x"),
		multipleFiles:     len(files) > 1,
	}

	var results []string

	for _, filename := range files {
		lines, err := readFile(filename)
		if err != nil {
			continue
		}
		matches := findPattern(pattern, lines, filename, flagsStruct)
		if flagsStruct.hasFileNames && len(matches) > 0 {
			results = append(results, filename)
		} else {
			results = append(results, matches...)
		}
	}
	return results
}
