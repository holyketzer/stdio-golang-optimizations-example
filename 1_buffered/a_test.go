package main

import (
	"os"
	"testing"
)

func BenchmarkSolveA(t *testing.B) {
	fin, _ := os.Open("../a.in")
	fout, _ := os.Create("../a.out")

	for i := 0; i < t.N; i++ {
		fin.Seek(0, 0)
		fout.Seek(0, 0)

		solve(fin, fout)
	}

	fin.Close()
	fout.Close()
}
