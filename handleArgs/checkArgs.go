package handleArgs

import (
	"errors"
	"path"

	"example.moh/handleFlag"
)

func CheckArgs(myArgs []string) (error, []string) {

	if len(myArgs) < 1 || len(myArgs) > 3 {
		return usageMessage(), []string{}
	}

	isValid, fileName := handleFlag.IsValidOutputFlag(myArgs[0])
	if len(myArgs) == 1 {
		if isValid {
			return usageMessage(), []string{}
		} else {
			return nil, myArgs
		}
	} else if len(myArgs) == 2 {
		if isValid {
			if len(fileName) < 1 {
				return usageMessage(), []string{}
			}
			myArgs = append(myArgs, "validFlag")
		} else {
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	} else if len(myArgs) == 3 {
		if isValid {
			myArgs = append(myArgs, "validFlag")
			if len(fileName) < 1 {
				return usageMessage(), []string{}
			}
			myArgs[2] = getBannerFileName(myArgs[2])
		}
	}

	return nil, myArgs
}

func usageMessage() error {
	// error strings should not be capitalized
	return errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func getBannerFileName(Banner string) string {
	if path.Ext(Banner) != ".txt" {
		Banner += ".txt"
	}
	return Banner
}
