package handleFlag

// // check if the flag is valid and return the file name
func IsValidOutputFlag(myFlag string) (bool, string) {
	isValid := false

	// fileExt := path.Ext(myFlag)

	outputFile := ""
	// handle out of range
	if len(myFlag) >= 8 && myFlag[:8] == "--output" {
		if len(myFlag) >= 9 {
			outputFile = myFlag[9:]
		}
		// if len(outputFile) >= 1 {
		isValid = true
		// }
	}
	return isValid, outputFile
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
//Usage: go run . [OPTION] [STRING]

//EX: go run . --color=<color> <letters to be colored> "something"
