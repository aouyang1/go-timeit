package timeit

import (
	"fmt"
	"time"
)

type TimeIt struct {
	start   time.Time
	oldMean float64
	newMean float64
	oldSS   float64
	newSS   float64
	iter    int
}

func New() TimeIt {
	return TimeIt{}
}

func (t *TimeIt) Reset() {
	t.start = time.Time{}
	t.oldMean = 0
	t.newMean = 0
	t.oldSS = 0
	t.newSS = 0
	t.iter = 0
}

func (t *TimeIt) Tic() {
	t.start = time.Now()
}

func (t *TimeIt) Toc() {
	t.iter++
	td := time.Now().Sub(t.start).Seconds()
	if t.iter == 1 {
		t.oldMean, t.newMean = td, td
	} else {
		t.newMean = t.oldMean + (td-t.oldMean)/float64(t.iter)
		t.newSS = t.oldSS + (td-t.oldMean)*(td-t.newMean)

		t.oldMean = t.newMean
		t.oldSS = t.newSS
	}
}

func (t TimeIt) Mean() float64 {
	if t.iter == 0 {
		return 0
	}
	return t.newMean
}

func (t TimeIt) StdDev() float64 {
	if t.iter <= 1 {
		return 0
	}
	return t.newSS / float64(t.iter-1)
}

func (t TimeIt) Show() {
	fmt.Printf("%.2f s ± %d ms per loop (mean ± std. dev. of %d runs, 1 loop each)\n", t.Mean(), int(t.StdDev()*1000), t.iter)
}
