package aoc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Object struct {
	line        int
	left, right int
}

type Field struct {
	Data [][]byte
}

type FieldCalcFunction func(f *Field) (int, int)

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

func (i *Field) Lines() []string {
	var lines []string
	for _, line := range i.Data {
		lines = append(lines, string(line))
	}
	return lines
}

func (i *Field) FindObjects(re string) []*Object {
	var objects []*Object
	matcher := regexp.MustCompile(re)

	for y, line := range i.Lines() {
		matches := matcher.FindAllStringIndex(line, -1)
		for _, match := range matches {
			objects = append(objects, &Object{
				line:  y,
				left:  match[0],
				right: match[1] - 1,
			})
		}
	}

	return objects
}

func (i *Field) ObjectAsString(o *Object) string {
	return string(i.Data[o.line][o.left : o.right+1])
}

func (i *Field) ObjectAsInt(o *Object) int {
	val, _ := strconv.Atoi(i.ObjectAsString(o))
	return val
}

func (o *Object) Adjacent(other *Object) bool {
	if other.right < o.left-1 || other.left > o.right+1 {
		return false
	}
	if other.line < o.line-1 || other.line > o.line+1 {
		return false
	}

	return true
}
