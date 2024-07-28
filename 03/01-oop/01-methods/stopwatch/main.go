package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	LapStart time.Time
	Results  []time.Duration
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()
	fmt.Println(sw.GetResults())

	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}

func (s *Stopwatch) Start() {
	for i := range s.Results {
		s.Results[i] = 0.0
	}
	s.Results = s.Results[len(s.Results):]
	s.LapStart = time.Now()
}

func (s *Stopwatch) SaveSplit() {
	s.Results = append(s.Results, time.Since(s.LapStart))
}

func (s *Stopwatch) GetResults() (d []time.Duration) {
	return s.Results
}
