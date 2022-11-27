package input

import (
	"fmt"

	"github.com/Radynsade/passgen/internal/ui/input/rules"
)

type YesNo struct {
	*rules.Set
	BaseInput
}

func NewYesNo(
	label string,
	placeholder string,
	required bool,
) *YesNo {
	rulesSet := rules.NewSet()
	rulesSet.AddRule(rules.OneOf("y", "n"))
	rulesSet.SetMessage("oneOf", "Type 'y' or 'n'.")

	return &YesNo{
		Set: rulesSet,
		BaseInput: BaseInput{
			Label:       label,
			Placeholder: placeholder,
			Required:    required,
		},
	}
}

func (yn *YesNo) PrintPlaceholder() {
	if yn.Required {
		print(yn.Placeholder + " (y/n)*: ")
	} else {
		print(yn.Placeholder + " (y/n): ")
	}
}

func (yn *YesNo) PrintAndRead() bool {
	yn.PrintLabel()

	var value string

	for {
		yn.PrintPlaceholder()

		fmt.Scanln(&value)
		validationMap, allValid := yn.Validate(value)

		if allValid {
			if value == "y" {
				return true
			}

			return false
		}

		for ruleName, valid := range validationMap {
			if valid {
				continue
			}

			yn.PrintError(ruleName)
		}
	}
}
