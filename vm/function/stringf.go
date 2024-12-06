package function

import (
	"fmt"
	"regexp"
	"strings"
	"vsce/vm/heap/cash"
)

func SprintF(input string, run bool) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			if strings.HasPrefix(match[1], "&") {
				if cash.Get_Local_Stack(run)[(match[1])[1:]] == nil {
					input = strings.Replace(input, fmt.Sprintf("{%v}", match[1]), "null", 1)
				} else {
					input = strings.Replace(input, fmt.Sprintf("{%v}", match[1]), fmt.Sprintf("%v", cash.Get_Local_Stack(run)[(match[1])[1:]].Value), 1)
				}
			} else {
				input = strings.Replace(input, fmt.Sprintf("{%v}", match[1]), match[1], 1)
			}
		}
	}
	return input
}
