package aoc

import (
	"bufio"
	"fmt"
	"os"
)

func Run(text string, fileName string, calcFunction FieldCalcFunction) {
	FieldFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
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
