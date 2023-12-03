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

	var data [][]byte

	scanner := bufio.NewScanner(FieldFile)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []byte(line))
	}

	if len(data[len(data)-1]) == 0 {
		data = data[:len(data)-1]
	}

	res1, res2 := calcFunction(&Field{Data: data})
	fmt.Printf("%s: %d, %d\n", text, res1, res2)
}
