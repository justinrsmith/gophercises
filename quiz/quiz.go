package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"log"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in format of question, answer")
	timeLimit := flag.Int("limit", 30, "time limit in seconds")
	flag.Parse()

	lines := openAndReadFile(*csvFilename)
	problems := parseLines(lines)

	runQuiz(problems, timeLimit)
}

func runQuiz(problems []problem, timeLimit *int) {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}
}

func readFile(reader io.Reader) ([][]string, error) {
    r := csv.NewReader(reader)
    lines, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    return lines, err
}

func openAndReadFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to read file: %s", fileName)
	}
	lines, err := readFile(file)
    if err != nil {
        fmt.Printf("Failed to read file: %s", fileName)
    }
    return lines
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer string
}