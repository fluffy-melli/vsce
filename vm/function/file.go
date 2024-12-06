package function

import (
	"fmt"
	"os"
)

func Read(f string) string {
	content, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("File Read Fail:", err)
		return ""
	}
	return string(content)
}
