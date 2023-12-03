package aoc

import (
	"bufio"
	"fmt"
	"os"

	"github.com/k0kubun/pp/v3"
)

func Run(text string, fileName string, calcFunction FieldCalcFunction) {
	_ = pp.Print // just to keep this module in the project

	FieldFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", text, err))
		return
	}

	var lines []string

	scanner := bufio.NewScanner(FieldFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	res1, res2 := calcFunction(NewField(lines))
	fmt.Printf("%s: %d, %d\n", text, res1, res2)
}
