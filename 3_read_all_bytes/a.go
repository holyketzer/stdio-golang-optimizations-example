package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
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

	offset := 0

	var n int
	var m int

	for buf[offset] != ' ' {
		n = n*10 + int(buf[offset]-'0')
		offset++
	}
	offset++

	for buf[offset] != 10 && buf[offset] != 13 {
		m = m*10 + int(buf[offset]-'0')
		offset++
	}

	var graph = make([][]int, n)

	for i := 0; i < m; i++ {
		var b int
		var e int

		for buf[offset] == 10 || buf[offset] == 13 {
			offset++
		}

		for buf[offset] != ' ' {
			b = b*10 + int(buf[offset]-'0')
			offset++
		}
		offset++

		for buf[offset] != 10 && buf[offset] != 13 {
			e = e*10 + int(buf[offset]-'0')
			offset++
		}

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
