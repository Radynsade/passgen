package export

import (
	"fmt"

	"golang.design/x/clipboard"
)

func ToClipboard(value string) {
	if err := clipboard.Init(); err != nil {
		fmt.Printf("Cannot write to clipboard: %s", err.Error())

		return
	}

	clipboard.Write(clipboard.FmtText, []byte(value))
}
