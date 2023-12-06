package main

import (
	_1201 "2023/1201"
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./inputs/01.txt"
	input, err := readInputFile(filePath)
	if err != nil {
		fmt.Printf("Error occured: %s\n", err.Error())
		return
	}

	result, err := _1201.Run01Two(input)
	if err != nil {
		fmt.Printf("Error occured: %s\n", err.Error())
		return
	}

	fmt.Printf("Result:\n%d\n", result)
}

func readInputFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
