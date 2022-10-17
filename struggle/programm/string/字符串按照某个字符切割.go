package string

import "strings"

func Split(s, str string) (result []string) {
	index := strings.Index(s, str)
	for index > -1 {
		result = append(result, s[:index])
		s = s[index+len(str):]
		index = strings.Index(s, str)
	}
	result = append(result, s)
	return
}
