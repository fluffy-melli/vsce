package function

import (
	"fmt"
	"regexp"
	"strings"
	"vsce/vm/heap/cash"
)

func SprintF(input string) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	matches := re.FindAllStringSubmatch(input[strings.LastIndex(input, "->")+2:], -1)
	if len(matches) > 0 {
		input = input[:strings.LastIndex(input, "->")]
		for _, match := range matches {
			if strings.HasPrefix(match[1], "&") {
				input = strings.Replace(input, "%v", fmt.Sprintf("%v", cash.Get_Local_Stack()[(match[1])[1:]].Value), 1)
			} else {
				input = strings.Replace(input, "%v", match[1], 1)
			}
		}
	}
	return input
}
