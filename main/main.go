package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	p "print/print"
)

const (
	Reset  = "\033[0m"
	red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// global slice for printing.
var result []string

func main() {
	colors := map[string]string{
		"red":    red,
		"green":  Green,
		"yellow": Yellow,
		"blue":   Blue,
		"purple": Purple,
		"cyan":   Cyan,
		"white":  White,
	}
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	count := 0
	arg1 := os.Args[1]
	// flag checking ...
	validcolor, color := IsValidColorFlag(arg1)
	validoutput, _ := IsValidOutputFlag(arg1)
	if len(os.Args) == 4 {
		if validcolor && len(arg1) >= 9 {

			file := os.Args[3]
			filename, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("error : file not found.")
				return
			}
			lines := strings.Split(string(filename), "\n")
			input := os.Args[2]
			words := strings.Split(input, "\\n")
			if len(input) == 0 {
				return
			}
			/// to display correctly in the file
			for a := 0; a < len(words); a++ {
				for i := 1; i < 9; i++ {
					endLine := false

					for _, char := range words[a] {
						s := int((char - 32) * 9)
						asciiLine := lines[s+i]
						result = append(result, colors[color]+asciiLine)
						endLine = true
					}
					if endLine {
						result = append(result, "\n")
					}
				}

				newLineCount := strings.Count(input, "\\n")
				// fmt.Print(words)
				if count < newLineCount && words[a] == "" {
					result = append(result, "\n")
					count++
				}
			}

			// print result
			for i := 0; i < len(result); i++ {
				fmt.Print(result[i])
			}
		}
	} else if len(os.Args) == 5 {
		if validcolor && len(arg1) >= 9 {
			file := os.Args[4]
			filename, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("error : file not found.")
				return
			}
			letterarg := os.Args[2]
			lines := strings.Split(string(filename), "\n")
			input := os.Args[3]
			words := strings.Split(input, "\\n")
			if len(input) == 0 {
				return
			}
			/// to display correctly in the file

			for a := 0; a < len(words); a++ {
				for i := 1; i < 9; i++ {
					endLine := false

					for _, char := range words[a] {
						s := int((char - 32) * 9)
						asciiLine := lines[s+i]
						if strings.ContainsRune(letterarg, char) {
							result = append(result, colors[color]+string(asciiLine)+Reset)
						} else {
							result = append(result, asciiLine)
						}

						endLine = true
					}
					if endLine {
						result = append(result, "\n")
					}
				}

				newLineCount := strings.Count(input, "\\n")
				// fmt.Print(words)
				if count < newLineCount && words[a] == "" {
					result = append(result, "\n")
					count++
				}
			}

			// print result
			for i := 0; i < len(result); i++ {
				fmt.Print(result[i])
			}
		}
	} else if len(os.Args) == 2 {
		if (validcolor || validoutput) && len(arg1) >= 9 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		} else if !validcolor {
			p.PrintLines("standard.txt")
		}
	} else if len(os.Args) == 3 {
		if (validcolor || validoutput) && len(arg1) >= 9 {
			file := "standard.txt"
			filename, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("error : file not found.")
				return
			}
			lines := strings.Split(string(filename), "\n")
			input := os.Args[2]
			words := strings.Split(input, "\\n")
			if len(input) == 0 {
				return
			}
			/// to display correctly in the file
			for a := 0; a < len(words); a++ {
				for i := 1; i < 9; i++ {
					endLine := false

					for _, char := range words[a] {
						s := int((char - 32) * 9)
						asciiLine := lines[s+i]
						result = append(result, colors[color]+asciiLine)
						endLine = true
					}
					if endLine {
						result = append(result, "\n")
					}
				}

				newLineCount := strings.Count(input, "\\n")
				// fmt.Print(words)
				if count < newLineCount && words[a] == "" {
					result = append(result, "\n")
					count++
				}
			}

			// print result
			for i := 0; i < len(result); i++ {
				fmt.Print(result[i])
			}

		} else if !validcolor {
			p.PrintLines(os.Args[2])
		}
	}
}

func IsValidColorFlag(myFlag string) (bool, string) {
	isValid := false

	// fileExt := path.Ext(myFlag)

	color := ""
	// handle out of range
	if len(myFlag) >= 7 && myFlag[:7] == "--color" {
		if len(myFlag) >= 8 {
			color = myFlag[8:]
		}
		// if len(outputFile) >= 1 {
		isValid = true
		// }
	}
	return isValid, color
}

func IsValidOutputFlag(myFlag string) (bool, string) {
	isValid := false

	// fileExt := path.Ext(myFlag)

	filename := ""
	// handle out of range
	if len(myFlag) >= 8 && myFlag[:8] == "--output" {
		if len(myFlag) >= 9 {
			filename = myFlag[9:]
		}
		// if len(outputFile) >= 1 {
		isValid = true
		// }
	}
	return isValid, filename
}

func getBannerFileName(Banner string) string {
	if path.Ext(Banner) != ".txt" {
		Banner += ".txt"
	}
	return Banner
}
