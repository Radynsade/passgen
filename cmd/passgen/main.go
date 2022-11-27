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

	if toFile := toFileInput.PrintAndRead(); toFile {
		export.ToFile(password)
	}
}

func buildLengthInput() *input.Number {
	length := input.NewNumber("", "Длина пароля", false)
	length.AddRule(rules.IsInteger())
	length.SetMin(3, true)
	length.SetMessage("isFloat", "Невозможно преобразовать в число.")
	length.SetMessage("isInteger", "Число должно быть целым.")
	length.SetMessage("minNumber", "Пароль не может быть меньше трёх символов.")

	return length
}

func buildUseLowercaseInput() *input.YesNo {
	lowercase := input.NewYesNo("", "Использовать латинские буквы нижнего регистра? (y/n)", false)
	lowercase.SetMessage("oneOf", "Введите y или n.")

	return lowercase
}

func buildUseUppercaseInput() *input.YesNo {
	uppercase := input.NewYesNo("", "Использовать латинские буквы верхнего регистра? (y/n)", false)
	uppercase.SetMessage("oneOf", "Введите y или n.")

	return uppercase
}

func buildUseNumbersInput() *input.YesNo {
	numbers := input.NewYesNo("", "Использовать цифры? (y/n)", false)
	numbers.SetMessage("oneOf", "Введите y или n.")

	return numbers
}

func buildUseSpecialsInput() *input.YesNo {
	specials := input.NewYesNo("", "Использовать специальные символы (. @ # $ % ...)? (y/n)", false)
	specials.SetMessage("oneOf", "Введите y или n.")

	return specials
}

func buildToFileInput() *input.YesNo {
	tofile := input.NewYesNo("", "Создать текстовой файл с паролем? (y/n)", false)
	tofile.SetMessage("oneOf", "Введите y или n.")

	return tofile
}
