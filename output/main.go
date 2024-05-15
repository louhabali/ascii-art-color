package main

import (
	"fmt"
	"os"
	"strings"

	"example.moh/getLines"
	"example.moh/handleFlag"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

/*
/////////////  TO DO:
- change the error msg when running :
-
*/
func main() {
	var result []string
	colors := []string{Red, Green, Yellow, Blue, Purple, Cyan, White}
	endLine := false
	count := 0

	lines, input := getLines.GetLines()
	words := strings.Split(input, "\\n")
	newLineCount := strings.Count(input, "\\n")

	if len(input) == 0 {
		return
	}
	/// to display correctly in the file
	result = append(result, "")
	for a := 0; a < len(words); a++ {
		for i := 1; i < 9; i++ {
			endLine = false

			for _, char := range words[a] {
				if int(char) < 32 || int(char) > 126 {
					fmt.Println("Error : char '", string(char), "' not found!!")
					return
				}
				s := (int(char) - 32) * 9

				asciiLine := lines[s+i]
				///for the third file
				asciiLine = strings.ReplaceAll(asciiLine, "\r", "")
				valid, color := handleFlag.IsValidColorFlag(os.Args[1])
				if valid && len(os.Args[1:]) >= 2 {
					for _, c := range colors {
						if color == c {
							result = append(result, c+asciiLine)
						}
					}
				}
				result = append(result, asciiLine)
				endLine = true
			}
			if endLine {
				result = append(result, "\n")
			}
		}

		if count < newLineCount && words[a] == "" {
			result = append(result, "\n")
			count++
		}

	}

	//////////////// O U T P U T ///////////////////
	valid, fileName := handleFlag.IsValidOutputFlag(os.Args[1])
	if !valid {
		// print result
		for i := 0; i < len(result); i++ {
			fmt.Print(result[i])
		}
	} else if len(os.Args[1:]) >= 2 && valid {
		writingErr := os.WriteFile(fileName, []byte(strings.Join(result, " ")), 0o644)
		////IF THERE IS AN ERROR WRITING THE FILE! EX :
		if writingErr != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}
	}
}
