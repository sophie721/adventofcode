package solver

import (
	"io"
	"os"
)

type Solver interface {
	SetInput(interface{}) error
	Solve(int) (interface{}, error)
}

func readFromInputFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	input := ""
	b := make([]byte, 1024)
	for {
		n, err := file.Read(b)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		input += string(b[:n])
	}
	return input, nil
}
