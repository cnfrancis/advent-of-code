package main

import "errors"

/*
Part I: CALIBRATION

Pass #1

Data structures
1. List of strings
2. Stack
3. Sum variable
4. Array size 2

Algo
*. Parse the string, push into stack if number else nothing
*. until end of string.
*. Pop item in array, put in last array position

** Edge cases (where algo will return incorrect data)
*. Only one 1 number in input
*. Zero number
*/
func parseString(calibrationLine string) (string, error) {
	if calibrationLine == "" {
		return "", errors.New("no input calibration input provided")
	}
	// Do logic
}
