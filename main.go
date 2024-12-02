package main

import (
	"fmt"
	"os"
	"strconv"

	"adventofcode/solver"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Invalid input: ", args)
		return
	}

	fmt.Printf("Day %s - %s\n", args[0], args[1])
	d, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Errorf("invalid input to parse as int: %s", d)
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Errorf("invalid input to parse as int: %s", n)
	}

	s, err := findSolver(d)
	if err != nil {
		panic(err)
	}
	output, err := s.Solve(n)
	if err != nil {
		fmt.Errorf("failed to solve day %d - %d: %w", d, n, err)
	}
	fmt.Println("output:")
	fmt.Println(output)
}

func findSolver(day int) (solver.Solver, error) {
	switch day {
	default:
		panic("not yet implemented: " + strconv.Itoa(day))
	case 1:
		return solver.NewSolver1("inputs/1.txt")
	}
}
