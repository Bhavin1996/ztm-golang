//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type Timeparse struct {
	msg   string
	input string
}

func (t *Timeparse) Error() string {
	return fmt.Sprintf("%v : %v", t.msg, t.input)
}

func parseTimne(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &Timeparse{"Invalid number of time componets", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &Timeparse{fmt.Sprintf("Error parsing hour %v", err), input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &Timeparse{fmt.Sprintf("Error parsing minute %v", err), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &Timeparse{fmt.Sprintf("Error parsing second %v", err), input}
		}
		if hour > 23 || hour < 0 {
			return Time{}, &Timeparse{"Hour value out of bounds", fmt.Sprintf("%v", hour)}
		}
		if minute > 60 || minute < 0 {
			return Time{}, &Timeparse{"Minute value out of bounds", fmt.Sprintf("%v", minute)}
		}
		if second > 60 || second < 0 {
			return Time{}, &Timeparse{"Minute value out of bounds", fmt.Sprintf("%v", second)}
		}
		return Time{hour, minute, second}, nil
	}
}
