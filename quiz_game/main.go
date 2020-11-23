package main

import (
  "flag"
  "fmt"
  "os"
  "encoding/csv"
  "strings"
  "time"
)

func main() {
  csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format of 'question,answer'")
  timeLimit := flag.Int("limit", 30, "Time limit per question in seconds.")
  flag.Parse()

  file, err := os.Open(*csvFilename)
  if err != nil {
    exit(fmt.Sprintf("Error: Failed to open the CSV file: %s\n", *csvFilename))
  }

  r := csv.NewReader(file)

  lines, err := r.ReadAll()

  if err != nil {
    exit(fmt.Sprintf("Error: Failed to parse the provided CSV file."))
  }

  problems := parseLines(lines)

  timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

  correct := 0

  for i, p := range problems {
    fmt.Printf("Problem #%d: %s = ", i+1, p.q)
    answerCh := make(chan string)
    go func() {
      var answer string
      fmt.Scanf("%s\n", &answer)
      answerCh <- answer
    }()

    select {
    case <- timer.C:
      fmt.Printf("\nYou got %d out of %d correct.\n", correct, len(problems))
      return
    case answer := <- answerCh:
      if answer == p.a {
        correct++
      }
    }
  }

  fmt.Printf("You got %d out of %d correct.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem{
  ret := make([]problem, len(lines))

  for i, line := range lines {
    ret[i] = problem{
      q: strings.TrimSpace(line[0]),
      a: strings.TrimSpace(line[1]),
    }
  }
  return ret
}

type problem struct {
  q string
  a string
}

func exit(msg string) {
  fmt.Println(msg)
  os.Exit(1)
}
