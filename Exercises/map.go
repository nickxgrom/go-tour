package Exercises

import (
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)

	m := make(map[string]int)

	for i := 0; i < len(arr); i++ {
		key := arr[i]

		var _, ok = m[key]

		if ok {
			m[key] += 1
		} else {
			m[key] = 1
		}
	}

	return m
}
