# Example of stdin/stdout optimization in competitive programming on Golang

## Example task: Connected components

Your task is to find connected components of a given undirected graph.

**Input**

The first line of the input file contains two integers 𝑛 and 𝑚 (1≤𝑛≤100000, 0≤𝑚≤200000) — the number of vertices and edges of the graph.

The following 𝑚 lines contain the descriptions of edges one per line. Edge number 𝑖 is described by two integers 𝑏𝑖, 𝑒𝑖 (1≤𝑏𝑖,𝑒𝑖≤𝑛) — the numbers of its ends. Loops and parallel edges are allowed.

**Output**

In the first line of the output file output an integer 𝑘, the number of connected components in the graph.

In the second line output 𝑛 numbers 𝑎1,𝑎2,…,𝑎𝑛 from 1 to 𝑘, where 𝑎𝑖 is the identifier of the connected component the 𝑖-th vertex belongs to.

**Examples**

input:
```
3 1
1 2
```

output:
```
2
1 1 2
```

input:
```
4 2
1 3
2 4
```

output:
```
2
1 2 1 2
```

## Solutions

Run benchmarks and generate CPU profiling graphs with:
```sh
go test -bench=. -cpuprofile cpu.prof && go tool pprof -svg cpu.prof > cpu.svg
```

Unbuffered IO:
```
$ go test -bench=. -cpuprofile cpu.prof && go tool pprof -svg cpu.prof > cpu.svg
goos: darwin
goarch: amd64
BenchmarkSolveA-8              1    3202774957 ns/op
PASS
ok      0_unbuffered 3.402s
```

Buffered IO:
```sh
$ go test -bench=. -cpuprofile cpu.prof && go tool pprof -svg cpu.prof > cpu.svg
goos: darwin
goarch: amd64
BenchmarkSolveA-8              4     285979620 ns/op
PASS
ok      1_buffered   2.456s
```

Read all input + convert to string and parse:
```sh
$ go test -bench=. -cpuprofile cpu.prof && go tool pprof -svg cpu.prof > cpu.svg
goos: darwin
goarch: amd64
BenchmarkSolveA-8             13      87102857 ns/op
PASS
ok      2_read_all_strings   2.346s
```

Read all input and parse raw bytes:
```sh
$ go test -bench=. -cpuprofile cpu.prof && go tool pprof -svg cpu.prof > cpu.svg
goos: darwin
goarch: amd64
BenchmarkSolveA-8             22      51912710 ns/op
PASS
ok      3_read_all_bytes 2.245s
```
