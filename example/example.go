package main

import (
	"time"

	"github.com/aouyang1/go-timeit"
)

func main() {
	t := timeit.New()
	for i := 0; i < 10; i++ {
		t.Tic()
		myFunc()
		t.Toc()
	}
	t.Show()
	// 0.24 s ± 0 ms per loop (mean ± std. dev. of 10 runs, 1 loop each)
}

func myFunc() {
	time.Sleep(234 * time.Millisecond)
}
