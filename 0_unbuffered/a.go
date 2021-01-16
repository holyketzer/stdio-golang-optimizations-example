package main

import (
	"fmt"
	"os"
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
	var n int
	var m int

	fmt.Fscan(fin, &n, &m)

	var graph = make([][]int, n)

	var b int
	var e int

	for i := 0; i < m; i++ {
		fmt.Fscan(fin, &b, &e)

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

	fmt.Fprintf(fout, "%d\n", max_color-1)
	for i := 0; i < n; i++ {
		fmt.Fprint(fout, colors[i], " ")
	}
	fmt.Fprintln(fout, "")
}

func main() {
	solve(os.Stdin, os.Stdout)
}
