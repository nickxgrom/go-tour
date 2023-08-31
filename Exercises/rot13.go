package Exercises

import (
	"io"
)

type Rot13Reader struct {
	R io.Reader
}

func (rot Rot13Reader) Read(p []byte) (int, error) {
	n, err := rot.R.Read(p)

	for i := 0; i < n; i++ {
		if (p[i] >= 65 && p[i] <= 77) || (p[i] > 97 && p[i] <= 109) {
			p[i] += 13
		} else if (p[i] > 77 && p[i] <= 90) || (p[i] > 109 && p[i] <= 122) {
			p[i] -= 13
		}
	}

	return n, err
}
