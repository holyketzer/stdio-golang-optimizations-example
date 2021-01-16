package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	DDN_FMT = "%d %d\n"
	DS_FMT  = "%d "
)

func dfs(v int, color int, graph *[][]int, colors *[]int) {
	(*colors)[v] = color

	for _, u := range (*graph)[v] {
		if (*colors)[u] != color {
			dfs(u, color, graph, colors)
		}
	}
}

func solve(fin *os.File, fout *os.File) {
	in := bufio.NewReader(fin)
	out := bufio.NewWriter(fout)

	var n int
	var m int

	fmt.Fscanf(in, DDN_FMT, &n, &m)

	var graph = make([][]int, n)

	var b int
	var e int

	for i := 0; i < m; i++ {
		fmt.Fscanf(in, DDN_FMT, &b, &e)

		b -= 1
		e -= 1

		graph[b] = append(graph[b], e)
		graph[e] = append(graph[e], b)
	}

	var colors = make([]int, n)
	var max_color = 1

	for i := 0; i < n; i++ {
		if colors[i] == 0 {
			dfs(i, max_color, &graph, &colors)
			max_color += 1
		}
	}

	fmt.Fprintf(out, "%d\n", max_color-1)
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, DS_FMT, colors[i])
	}
	fmt.Fprintf(out, "\n")
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
