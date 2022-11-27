package rules

import "strconv"

var (
	numbers = []rune{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
	}
)

// Check if string can be converted to any number type except complex.
func isFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)

	return err == nil
}

func isInteger(value string) bool {
	_, err := strconv.Atoi(value)

	return err == nil
}

func maxNumber(
	value float64,
	max float64,
	included bool,
) bool {
	if included {
		return value <= max
	}

	return value < max
}

func minNumber(
	value float64,
	min float64,
	included bool,
) bool {
	if included {
		return value >= min
	}

	return value > min
}
