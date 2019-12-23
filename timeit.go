package timeit

import (
	"fmt"
	"time"
)

// TimeIt keeps track of the running mean and sum of squares for a series of time measurements
type TimeIt struct {
	start   time.Time
	oldMean float64
	newMean float64
	oldSS   float64
	newSS   float64
	iter    int
}

// New creates a new TimeIt struct that will hold state of all time measurements
func New() TimeIt {
	return TimeIt{}
}

// Reset resets the timit struct to be used again without having to reallocate a new struct
func (t *TimeIt) Reset() {
	t.start = time.Time{}
	t.oldMean = 0
	t.newMean = 0
	t.oldSS = 0
	t.newSS = 0
	t.iter = 0
}

// Tic starts the timer for the current measurement
func (t *TimeIt) Tic() {
	t.start = time.Now()
}

// Toc adds a time delta from the last tic to the timeit summary
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

func (t TimeIt) mean() float64 {
	if t.iter == 0 {
		return 0
	}
	return t.newMean
}

func (t TimeIt) stdDev() float64 {
	if t.iter <= 1 {
		return 0
	}
	return t.newSS / float64(t.iter-1)
}

// Show prints out the current summary of the recorded timings
func (t TimeIt) Show() {
	fmt.Printf("%s ± %s per loop (mean ± std. dev. of %d runs, 1 loop each)\n", printTime(t.mean()), printTime(t.stdDev()), t.iter)
}

func printTime(t float64) string {
	if t < 1 {
		return fmt.Sprintf("%d ms", int(t*1000))
	}
	return fmt.Sprintf("%.2f s", t)
}
