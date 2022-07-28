package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type lines struct {
	left  int
	right int
	count int
}

func main() {
	after := flag.Int("A", 2, "after key")
	before := flag.Int("B", 0, "before key")
	context := flag.Int("C", 0, "context key")
	count := flag.Int("c", 0, "count key")
	ignoreCase := flag.Bool("i", false, "ignore-case key")
	inverter := flag.Bool("v", false, "invertor key")
	fixed := flag.Bool("fixed", false, "fixed key")
	lineNum := flag.Bool("n", false, "live num key")
	flag.Parse()

	l := lines{0, 0, 0}
	file := readFile("input.txt")
	t := "апнутся"
	if *count != 0 {
		l.count = *count
	} else {
		l.count = len(file)
	}

	switch {
	case *context != 0:
		l.left, l.right = *context, *context
	case *after != 0:
		l.right = *after
	case *before != 0:
		l.left = *before
	default:
		l.right = len(file)
	}

	for i, value := range file {
		if processLine(value, t, *ignoreCase, *inverter, *fixed) {
			max := math.Max(0, float64(i-l.left))
			min := math.Min(float64(len(file)-1), float64(i+l.right))

			for j := max; j <= min; j++ {
				if l.count >= 1 {
					if *lineNum {
						fmt.Println("num: ", j, "line: ", file[int(j)])
						l.count--
					} else {
						fmt.Println("line: ", file[int(j)])
						l.count--
					}
				} else {
					os.Exit(1)
				}
			}
		}
	}
}

func readFile(name string) []string {
	file, _ := os.Open(name)
	defer file.Close()
	byteFile, _ := io.ReadAll(file)
	return strings.Split(string(byteFile), "\n")
}

func processLine(str string, temp string, ignoreCase bool, invert bool, fixed bool) bool {
	if ignoreCase {
		str = strings.ToLower(str)
		temp = strings.ToLower(temp)
	}
	if fixed {
		if str == temp {
			return true && !invert
		}
		return false || invert
	}

	if strings.Contains(str, temp) {
		return true && !invert
	}
	return false || invert
}
