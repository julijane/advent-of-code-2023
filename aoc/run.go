package aoc

import (
	"bufio"
	"fmt"
	"os"

	"github.com/k0kubun/pp/v3"
)

// Run runs the given calcFunction on the given input file
func Run(text string, fileName string, calcFunction InputCalcFunction) {
	_ = pp.Print // just to keep this module in the project

	InputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", text, err))
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

	res1, res2 := calcFunction(NewInput(lines))
	fmt.Printf("%s: %d, %d\n", text, res1, res2)
}
