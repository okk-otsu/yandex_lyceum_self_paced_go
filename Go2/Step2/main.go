package main

import (
	"bufio"
	"errors"
	"os"
	"time"
)

var ErrNoLinesInRange = errors.New("no log lines in range")

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	f, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var result []string

	for sc.Scan() {
		line := sc.Text()
		if len(line) < 10 {
			continue
		}

		datePart := line[:10]
		date, err := time.Parse("02.01.2006", datePart)
		if err != nil {
			continue
		}

		if (date.After(start) || date.Equal(start)) && (date.Before(end) || date.Equal(end)) {
			result = append(result, line)
		}
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, ErrNoLinesInRange
	}

	return result, nil
}
