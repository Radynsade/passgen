package gen

import (
	"math/rand"
	"strings"
	"time"
)

func Generate(options *Options) string {
	symbols := []rune{}

	if options.UseLowercase {
		symbols = append(symbols, lowercase...)
	}

	if options.UseUppercase {
		symbols = append(symbols, uppercase...)
	}

	if options.UseNumbers {
		symbols = append(symbols, numbers...)
	}

	if options.UseSpecials {
		symbols = append(symbols, specials...)
	}

	if len(symbols) == 0 {
		symbols = append(symbols, lowercase...)
	}

	min := 0
	max := len(symbols) - 1
	var sb strings.Builder

	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= int(options.Length); i++ {
		sb.WriteRune(symbols[randomInt(min, max)])
	}

	return sb.String()
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}
