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

type Timer struct {
	Hour   int
	Minute int
	Second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v:%v", t.msg, t.input)
}

func ParseTime(time string) (Timer, error) {
	if time == "" {
		return Timer{}, &TimeParseError{
			msg: "Invalid input",
		}
	}

	timeSlice := strings.Split(time, ":")

	if len(timeSlice) != 3 {
		return Timer{}, &TimeParseError{
			msg: "Invalid input length",
		}
	}

	h, err := strconv.Atoi(timeSlice[0])

	if err != nil {
		return Timer{}, &TimeParseError{
			msg:   "Invalid hour value",
			input: timeSlice[0],
		}
	}

	m, err := strconv.Atoi(timeSlice[1])

	if err != nil {
		return Timer{}, &TimeParseError{
			msg:   "Invalid minute value",
			input: timeSlice[1],
		}
	}

	s, err := strconv.Atoi(timeSlice[2])

	if err != nil {
		return Timer{}, &TimeParseError{
			msg:   "Invalid second value",
			input: timeSlice[2],
		}
	}

	if h > 23 || h < 0 {
		return Timer{}, &TimeParseError{
			msg:   "hour out of range, range must be between 23 and 0",
			input: fmt.Sprintf("%d", h),
		}
	}

	if m > 59 || m < 0 {
		return Timer{}, &TimeParseError{
			msg:   "minute out of range, range must be between 59 and 0",
			input: fmt.Sprintf("%d", m),
		}
	}

	if s > 59 || s < 0 {
		return Timer{}, &TimeParseError{
			msg:   "second out of range, range must be between 59 and 0",
			input: fmt.Sprintf("%d", s),
		}
	}

	result := Timer{
		Hour:   h,
		Minute: m,
		Second: s,
	}

	return result, nil

}
