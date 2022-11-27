package main

import (
	"github.com/Radynsade/passgen/internal/export"
	"github.com/Radynsade/passgen/internal/gen"
	"github.com/Radynsade/passgen/internal/ui/input"
	"github.com/Radynsade/passgen/internal/ui/input/rules"
)

var (
	lengthInput       = buildLengthInput()
	useLowercaseInput = buildUseLowercaseInput()
	useUppercaseInput = buildUseUppercaseInput()
	useNumbersInput   = buildUseNumbersInput()
	useSpecialsInput  = buildUseSpecialsInput()
	toFileInput       = buildToFileInput()
	toClipboardInput  = buildToClipboardInput()
)

func main() {
	password := gen.Generate(gen.NewOptions(
		uint(lengthInput.PrintAndRead()),
		useLowercaseInput.PrintAndRead(),
		useUppercaseInput.PrintAndRead(),
		useNumbersInput.PrintAndRead(),
		useSpecialsInput.PrintAndRead(),
	))

	println(password)

	if toClipboard := toClipboardInput.PrintAndRead(); toClipboard {
		export.ToClipboard(password)
	}

	if toFile := toFileInput.PrintAndRead(); toFile {
		export.ToFile(password)
	}
}

func buildLengthInput() *input.Number {
	length := input.NewNumber("", "Password length", false)
	length.AddRule(rules.IsInteger())
	length.SetMin(3, true)
	length.SetMessage("isFloat", "It is not a number.")
	length.SetMessage("isInteger", "Must be an integer.")
	length.SetMessage("minNumber", "Cannot be less than 3 characters.")

	return length
}

func buildUseLowercaseInput() *input.YesNo {
	lowercase := input.NewYesNo("", "Use latin lowercase letters?", false)

	return lowercase
}

func buildUseUppercaseInput() *input.YesNo {
	uppercase := input.NewYesNo("", "Use latin uppercase letters?", false)

	return uppercase
}

func buildUseNumbersInput() *input.YesNo {
	numbers := input.NewYesNo("", "Use numbers?", false)

	return numbers
}

func buildUseSpecialsInput() *input.YesNo {
	specials := input.NewYesNo("", "Use special characters? (. @ # $ % ...)?", false)

	return specials
}

func buildToFileInput() *input.YesNo {
	tofile := input.NewYesNo("", "Create text file with password?", false)

	return tofile
}

func buildToClipboardInput() *input.YesNo {
	tofile := input.NewYesNo("", "Save password to clipboard?", false)

	return tofile
}
