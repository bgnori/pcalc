package main

import (
	"encoding/json"
	"flag"
	. "github.com/bgnori/npoker"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

func calc(fin *os.File, fout *os.File, trials int) {
	var req Request

	decoder := json.NewDecoder(fin)
	decoder.Decode(&req)

	/*
		if req.Goroutines > 1 {
			runtime.GOMAXPROCS(runtime.NumCPU())
		}
	*/

	if trials > 0 {
		req.Trials = trials
	}

	if req.Seed == nil {
		req.Seed = GetSeedFromRAND()
	}

	r := NewRand()
	r.SeedFromBytes(req.Seed)

	summary := NewSummary(req)
	summary.Start = time.Now()

	w := NewWorkSet(req.Board, req.Players)
	for i := 0; i < req.Trials; i++ {
		u := w.Clone()
		summary.Add(u.Run(r))
	}
	summary.End = time.Now()

	encoder := json.NewEncoder(fout)
	encoder.Encode(summary)
}

func calc2(fin *os.File, fout *os.File, trials int) {
	var req Request

	decoder := json.NewDecoder(fin)
	decoder.Decode(&req)
	summary := NewSummary(req)
	summary.Start = time.Now()

	w := NewWorkSet(req.Board, req.Players)
	w.ByComb(summary)
	summary.End = time.Now()

	encoder := json.NewEncoder(fout)
	encoder.Encode(summary)
}

func main() {
	var count int
	var input string
	var output string
	var spprof string
	var fin *os.File
	var fout *os.File
	var err error
	var auto bool

	flag.IntVar(&count, "c", 0, "count for experiment")
	flag.BoolVar(&auto, "a", false, "auto filename based on input, *request.json => *result.json")
	flag.StringVar(&input, "in", "", "input file. default is STDIN")
	flag.StringVar(&output, "out", "", "output file. default is STDOUT")
	flag.StringVar(&spprof, "pprof", "", "pprof file.")
	flag.Parse()

	if len(input) == 0 {
		fin = os.Stdin
		input = "STDIN"
	} else {
		fin, err = os.Open(input)
		if err != nil {
			panic(err)
		}
		defer fin.Close()
	}

	if auto {
		output = strings.Replace(input, "request.json", "result.json", 1)
	}

	if len(output) == 0 {
		fout = os.Stdout
	} else {
		fout, err = os.Create(output)
		if err != nil {
			if auto {
				return
			} else {
				panic(err)
			}
		}
		defer fout.Close()
	}

	if len(spprof) > 0 {
		var fpprof *os.File
		fpprof, err := os.Create(spprof)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
		pprof.StartCPUProfile(fpprof)
	}

	//calc(fin, fout, count)
	calc2(fin, fout, count)
}
