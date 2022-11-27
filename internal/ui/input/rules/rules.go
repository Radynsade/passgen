package rules

import (
	"regexp"
	"strconv"
)

var (
	emailRegexp, _ = regexp.Compile("^(([^<>()[\\]\\.,;:\\s@\"]+(\\.[^<>()[\\]\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$")
)

// First return value is the name of the rule, second is result (valid or not).
type Rule func(value string) (string, bool)

func IsFloat() Rule {
	const ruleName = "isFloat"

	return func(value string) (string, bool) {
		return ruleName, isFloat(value)
	}
}

func IsInteger() Rule {
	const ruleName = "isInteger"

	return func(value string) (string, bool) {
		isDigit := false

		for _, char := range value {
			for _, number := range numbers {
				if char == number {
					isDigit = true
					break
				}
			}

			if !isDigit {
				return ruleName, false
			}

			isDigit = false
		}

		return ruleName, true
	}
}

func MaxNumber(
	max float64,
	included bool,
) Rule {
	const ruleName = "maxNumber"

	return func(value string) (string, bool) {
		num, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return ruleName, false
		}

		return ruleName, maxNumber(num, max, included)
	}
}

func MinNumber(
	min float64,
	included bool,
) Rule {
	const ruleName = "minNumber"

	return func(value string) (string, bool) {
		num, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return ruleName, false
		}

		return ruleName, minNumber(num, min, included)
	}
}

func OneOf(values ...string) Rule {
	const ruleName = "oneOf"

	return func(value string) (string, bool) {
		for _, v := range values {
			if v == value {
				return ruleName, true
			}
		}

		return ruleName, false
	}
}
