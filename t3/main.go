package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	u := flag.Bool("u", false, "skip duplicate lines")
	n := flag.Bool("n", false, "sort by numeric")
	r := flag.Bool("r", false, "reverse")
	k := flag.Int("k", 0, "column")
	output := flag.Bool("o", false, "write sort result in file, or stdout(default-file)")
	content := fetchFileData("")
	switch {
	case *u:
		content = unique(content)
	case *n:
		content = sortByNumeric(content)
	case *k != 0:
		content = sortByColumn(content, *k)
	case *r:
		content = reverse(content)
	}

	if *output {
		fmt.Println(content)
	} else {
		writeResult("out.txt", content)
	}

}

func reverse(content []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(content)))
	return content
}

func sortByColumn(content []string, column int) []string {
	sort.Slice(content, func(i, j int) bool {
		return content[i][column] < content[j][column]
	})

	return content
}

func sortByNumeric(content []string) []string {
	for i := range content {
		var c int
		splited := strings.Split(content[i], " ")
		for m := range splited {
			_, err := strconv.Atoi(splited[m])
			if err != nil {
				c++
			}
		}

		if c == len(splited) {
			var res []int
			for m := range splited {
				k, _ := strconv.Atoi(splited[m])
				res = append(res, k)
			}

			sort.Ints(res)
			for m := range res {
				splited = append(splited[:m], strconv.Itoa(res[m]))
			}

		}
		content[i] = strings.Join(splited, " ")
	}
	return content
}

func unique(content []string) []string {
	for i, str := range content {
		for j := i + 1; j < len(content); j++ {
			if str == content[j] {
				content = append(content[:i], content[j:]...)
			}
		}
	}
	return content
}

func writeResult(path string, content []string) {
	out, _ := os.Create(path)
	defer out.Close()

	for _, v := range content {
		out.WriteString(v + "\n")
	}
}

func fetchFileData(path string) (result []string) {
	file, _ := os.Open(path)
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		row := sc.Text()
		result = append(result, row)
	}
	return
}
