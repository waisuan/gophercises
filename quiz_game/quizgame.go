package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	results = map[string]int{"correct": 0, "incorrect": 0}
)

func main() {
	rawTimer := flag.Int("timer", 30, "set timer for quiz (in seconds)")
	problemsFile := flag.String("file", "problems.csv", "path to file containing list of problem questions")
	flag.Parse()

	fmt.Println(fmt.Sprintf("Will fetch problems from %s", *problemsFile))

	f, err := os.Open(*problemsFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to open %s: %s", *problemsFile, err))
	}
	defer f.Close()

	problems, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to fetch problems: %s", err))
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press <ENTER> to start the quiz.")
	scanner.Scan()

	fmt.Println(fmt.Sprintf("Quiz will run for %d seconds.", *rawTimer))
	done := make(chan bool)
	go quiz(done, problems)

	timer := time.Duration(*rawTimer) * time.Second
	select {
		case <- done:
			printResults(len(problems))
		case <- time.After(timer):
			fmt.Println("!!! Times's up !!!")
			printResults(len(problems))
	}
}

func printResults(problemSize int) {
	fmt.Println(fmt.Sprintf("Results: %d/%d", results["correct"], problemSize))
}

func quiz(ch chan<- bool, problems [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	for _, problem := range problems {
		fmt.Println(fmt.Sprintf(">>> %s", problem[0]))
		scanner.Scan()
		if scanner.Text() == problem[1] {
			results["correct"]++
		} else {
			results["incorrect"]++
		}
	}

	ch <- true
}
