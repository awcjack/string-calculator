package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// func

func StringAdd(input string) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}

	separator := ""
	if len(input) > 3 {
		if string(input[0]) == "/" && string(input[1]) == "/" {
			separatorEnd := strings.Index(input, "\n")
			if separatorEnd == -1 {
				return 0, errors.New("missing separator end newline character")
			}
			if separatorEnd > 3 {
				count := strings.Count(input, "[")
				if count > 1 {
					for i := 0; i < count; i++ {
						arbitrarySeparatorEnd := 0
						for j := arbitrarySeparatorEnd; j < separatorEnd-3+i*3; j++ {
							if string(input[3+i*3+j]) == "]" {
								arbitrarySeparatorEnd = j - arbitrarySeparatorEnd
								break
							}
						}
						input = strings.ReplaceAll(input, string(input[3+i*3:3+i*3+arbitrarySeparatorEnd]), ",")
						if arbitrarySeparatorEnd > 1 {
							// replace removed some character
							separatorEnd = strings.Index(input, "\n")
						}
					}
				} else {
					separator = string(input[3 : separatorEnd-1])
				}
			} else {
				separator = string(input[2])
			}
			input = input[separatorEnd:]
		}
		if separator != "" {
			input = strings.ReplaceAll(input, separator, ",")
		}
	}
	if !strings.Contains(input, ",") && !strings.Contains(input, "\n") && !strings.Contains(input, separator) {
		result, err := strconv.Atoi(input)
		if err != nil {
			return 0, err
		}
		return result, nil
	}

	input = strings.ReplaceAll(input, "\n", ",")
	inputArray := strings.Split(input, ",")
	numberArray := make([]int, len(inputArray))
	negativeNumberArray := make([]string, 0)
	for index, numberString := range inputArray {

		if numberString != "" {
			if string(numberString[0]) == "-" {
				negativeNumberArray = append(negativeNumberArray, numberString)
			} else {
				number, err := strconv.Atoi(numberString)
				if err != nil {
					return 0, err
				}
				if number > 1000 {
					numberArray[index] = 0
				} else {
					numberArray[index] = number
				}
			}
		}
	}
	if len(negativeNumberArray) != 0 {
		return 0, fmt.Errorf("negatives not allowed: %s", strings.Join(negativeNumberArray, " "))
	}

	sum := 0
	for _, number := range numberArray {
		sum += number
	}

	return sum, nil
}
