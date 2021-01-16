package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	buf, _ := ioutil.ReadAll(fin)
	out := bufio.NewWriter(fout)

	str := string(buf)
	lines := strings.Split(str, "\n")

	line0 := strings.Split(lines[0], " ")
	n, _ := strconv.Atoi(line0[0])
	m, _ := strconv.Atoi(line0[1])

	var graph = make([][]int, n)

	for i := 1; i < m; i++ {
		line := strings.Split(lines[i], " ")

		b, _ := strconv.Atoi(line[0])
		e, _ := strconv.Atoi(line[1])

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

	out.WriteString(strconv.Itoa(max_color - 1))
	out.WriteByte('\n')

	for i := 0; i < n; i++ {
		out.WriteString(strconv.Itoa(colors[i]))
		out.WriteByte(' ')
	}

	out.WriteByte('\n')
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
