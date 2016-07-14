package main

import (
	"encoding/json"
	"fmt"
	. "github.com/bgnori/npoker"
	"os"
	//"runtime"
	"time"
)

/*
type Request struct {
	Source     string `json:"source"`
	Players    []Deck `json:"players"`
	Board      []Deck `json:"board"`
	Trials     int    `json:"trials"`
	Goroutines int    `json:"goroutines"`
}
*/

func main() {

	var req Request
	decoder := json.NewDecoder(os.Stdin)
	decoder.Decode(&req)

	fmt.Printf("%+v\n", req)

	/*
		if req.Goroutines > 1 {
			runtime.GOMAXPROCS(runtime.NumCPU())
		}
	*/

	w := NewWorkSet(req.Board, req.Players)

	summary := NewSummary(req.Players)
	r := NewRand()
	b := GetSeedFromRAND()
	r.SeedFromBytes(b)

	start := time.Now()
	for i := 0; i < req.Trials; i++ {
		u := w.Clone()
		summary.Add(u.Run(r))
	}
	end := time.Now()

	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	fmt.Println(summary)
}
