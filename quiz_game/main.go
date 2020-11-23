package main

import (
  "flag"
  "fmt"
  "os"
  "encoding/csv"
  "strings"
)

func main() {
  csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format of 'question,answer'")
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
  fmt.Println(problems)

  correct := 0

  for i, p := range problems {
    fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
    var answer string
    fmt.Scanf("%s\n", &answer)
    if answer == p.a {
      correct++
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
