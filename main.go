package main

import (
	"encoding/json"
	//"fmt"
	. "github.com/bgnori/npoker"
	"os"
	//"runtime"
	"time"
)

func main() {

	var req Request
	decoder := json.NewDecoder(os.Stdin)
	decoder.Decode(&req)

	//fmt.Printf("%+v\n", req)

	/*
		if req.Goroutines > 1 {
			runtime.GOMAXPROCS(runtime.NumCPU())
		}
	*/

	w := NewWorkSet(req.Board, req.Players)

	r := NewRand()
	b := GetSeedFromRAND()
	summary := NewSummary(req, b)
	r.SeedFromBytes(b)

	summary.Start = time.Now()
	for i := 0; i < req.Trials; i++ {
		u := w.Clone()
		summary.Add(u.Run(r))
	}
	summary.End = time.Now()

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(summary)
}
