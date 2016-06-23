package main

import (
	"encoding/json"
	"fmt"
	. "github.com/bgnori/npoker"
	"os"
	"runtime"
	"time"
)

type Request struct {
	Source     string `json:"source"`
	Players    []Deck `json:"players"`
	Board      []Deck `json:"board"`
	Trials     int    `json:"trials"`
	Goroutines int    `json:"goroutines"`
}

func main() {

	var r Request
	decoder := json.NewDecoder(os.Stdin)
	decoder.Decode(&r)

	fmt.Printf("%+v\n", r)

	if r.Goroutines > 1 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	calc := NewEqCalc(r.Board, r.Players)

	ex := NewRunner(r.Trials, r.Goroutines,
		calc,
		NewEqSummarizer(calc),
	)

	start := time.Now()
	ex.Run()
	end := time.Now()

	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	fmt.Println(ex.Summary())
}
