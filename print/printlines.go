package print

import (
	"fmt"
	"os"
	"strings"
)

func PrintLines(argy string) {
	var result []string
	endLine := false
	count := 0

	input := os.Args[1]

	file, err := os.ReadFile(argy)
	if err != nil {
		fmt.Println("error : file not found.")
		return
	}
	lines := strings.Split(string(file), "\n")
	words := strings.Split(input, "\\n")
	newLineCount := strings.Count(input, "\\n")

	if len(input) == 0 {
		return
	}
	/// to display correctly in the file
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
				// for thinkertoy file
				asciiLine = strings.ReplaceAll(asciiLine, "\r", "")

				result = append(result, asciiLine)

				endLine = true
			}
			if endLine {
				result = append(result, "\n")
			}
		}
		// Check if there are newlines to be inserted
		if count < newLineCount {
			result = append(result, "\n")
			count++
		}
	}
	// Print the result
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}
}
