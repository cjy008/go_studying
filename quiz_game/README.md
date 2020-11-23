## Quiz Game
This program will read a quiz provided via a CSV file and will then conduct the quiz to the user. The program will track the number of correct and incorrect answers. Regardless of whether the answer was correct or incorrect, the following questions will be asked.

```
Usage of ./quiz_game:
  -csv string
    	A CSV file in the format of 'question,answer' (default "problems.csv")
  -limit int
    	Time limit per question in seconds. (default 30)
```

The CSV file is defaulted to `problems.csv` but is customizable by the user with the `-csv` flag. The format of the CSV file is question in the first column and answer in the second one in the same row.

The time limit on the quiz is defaulted to 30 seconds. It can be modified via the `-limit` flag. The quiz will end after the time limit has been exceeded. It will end the quiz and not wait for any final answers or interactions from the user once the time limit is exceeded. The answer has to be entered before the time limit expires to be counted.

An assumption made is that the CSV file will be relatively small (<100 lines) and the answers will be single words or numbers.

At the end of the quiz, a final score will be outputted of total correct out of the total number of questions. Invalid answers are counted as incorrect answers.


### Build and Execute
```
go build . && ./quiz_game
```
