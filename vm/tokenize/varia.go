package tokenize

import (
	"fmt"
	"regexp"
)

func GET_VAR(line string) {
	re := regexp.MustCompile(`&[^,)\s]+`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		fmt.Println(match[1:])
	}
}
