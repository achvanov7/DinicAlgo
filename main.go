package main

import (
	"DinicAlgo/dinic"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	tmp := GetInputSlice(scanner)
	n, m := tmp[0], tmp[1]
	dinicAlgo := new(dinic.Dinic)
	dinicAlgo.Init(n, 0, n-1)
	for i := 0; i < m; i++ {
		tmp = GetInputSlice(scanner)
		u, v, c := tmp[0], tmp[1], tmp[2]
		dinicAlgo.AddEdge(u-1, v-1, c, 0)
	}
	res := dinicAlgo.MaxFlow()
	os.WriteFile("output.txt", []byte(strconv.FormatInt(int64(res), 10)), 0644)
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func GetInputSlice(scanner *bufio.Scanner) []int {
	scanner.Scan()
	return numbers(scanner.Text())
}
