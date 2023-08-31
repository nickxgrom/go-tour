//https://go.dev/tour/methods/22

package Exercises

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}
