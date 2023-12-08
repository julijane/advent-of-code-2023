package aoc

import (
	"bufio"
	"fmt"
	"os"

	"github.com/k0kubun/pp/v3"
)

// InputCalcFunction is the function signature for the calculation function
type InputCalcFunction func(i *Input, runPart1, runpart2 bool) (int, int)

// Run runs the given calcFunction on the given input file
func Run(fileName string, calcFunction InputCalcFunction, runPart1, runPart2 bool) {
	_ = pp.Print // just to keep this module in the project

	InputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", fileName, err))
		return
	}

	var lines []string

	scanner := bufio.NewScanner(InputFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	res1, res2 := calcFunction(NewInput(lines), runPart1, runPart2)

	result := fileName + ":"
	if runPart1 {
		result += fmt.Sprintf(" part 1: %d", res1)
	}
	if runPart2 {
		if runPart1 {
			result += ", "
		}
		result += fmt.Sprintf(" part 2: %d", res2)
	}

	fmt.Println(result)
}
