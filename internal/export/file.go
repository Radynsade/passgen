package export

import (
	"fmt"
	"os"
	"time"
)

func ToFile(value string) {
	timestamp := time.Now().UnixNano()
	filename := "passgen_" + fmt.Sprintf("%d", timestamp) + ".txt"

	err := os.WriteFile(filename, []byte(value), 0777)

	if err != nil {
		panic(err)
	}
}
