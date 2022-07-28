package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var fields = flag.Int("f", -1, "выбрать поля (колонки)")
	var delimiter = flag.String("d", "\t", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "только строки с разделителем")
	log.SetFlags(0)
	flag.Parse()

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	delim := []byte(*delimiter)
	lines := bytes.Split(content, []byte("\n"))

	if *separated {
		separate(lines, delim)
	}
	delimit(lines, delim, fields)
}

func delimit(lines [][]byte, delimiter []byte, fields *int) {
	res := make([][][]byte, 0)
	for i := range lines {
		temp := bytes.Split(lines[i], delimiter)
		if field := *fields; field > -1 {
			if len(temp) > field {
				temp = [][]byte{temp[field]}
			} else {
				temp = [][]byte{}
			}
		}
		res = append(res, temp)
	}

	fmt.Printf("%q", res)
}

func separate(lines [][]byte, delimiter []byte) {
	temp := make([][]byte, 0)
	for i := range lines {
		if bytes.Contains(lines[i], delimiter) {
			temp = append(temp, lines[i])
		}
	}
	lines = temp

	fmt.Printf("%q", lines)
}
