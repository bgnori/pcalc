package main

import (
	"flag"
	"fmt"
	. "github.com/bgnori/npoker"
	"runtime"
	"time"
)

func main() {
	//source: https://twitter.com/PokerStarsJapan/status/725880733093400576
	/*
		PlayerOneHole := Deck{
			Card{NINE, DIAMONDS},
			Card{SIX, DIAMONDS},
		}

		PlayerTwoHole := Deck{
			Card{FIVE, CLUBS},
			Card{FIVE, HEARTS},
		}

		PlayerThreeHole := Deck{
			Card{JACK, CLUBS},
			Card{EIGHT, CLUBS},
		}

		board := Deck{
			Card{NINE, CLUBS},
			Card{SEVEN, CLUBS},
			Card{FIVE, DIAMONDS},
			Card{FOUR, DIAMONDS},
		}
	*/

	nexp := flag.Int("exp", 100, "the number of experiments")
	nroutine := flag.Int("rtn", 1, "the number of  go routine to calculate board")
	flag.Parse()

	PlayerOneHole := Deck{
		Card{JACK, SPADES},
		Card{JACK, CLUBS},
	}

	PlayerTwoHole := Deck{
		Card{ACE, SPADES},
		Card{ACE, HEARTS},
	}

	PlayerThreeHole := Deck{
		Card{KING, DIAMONDS},
		Card{ACE, DIAMONDS},
	}
	board := Deck{}

	fmt.Printf("exp=%d, worker=%d\n", *nexp, *nroutine)
	fmt.Printf("NumCPU=%d, GOMAXPROCS=%d\n", runtime.NumCPU(), runtime.GOMAXPROCS(0))
	if *nexp > 1 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	fmt.Printf("%v\n", board)
	calc := NewEqCalc(board, []Deck{PlayerOneHole, PlayerTwoHole, PlayerThreeHole})

	ex := NewRunner(*nexp, *nroutine,
		calc,
		NewEqSummarizer(calc),
	)

	start := time.Now()
	ex.Run()
	end := time.Now()

	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	fmt.Println(ex.Summary())

	/*
		players := []Deck{PlayerOneHole, PlayerTwoHole, PlayerThreeHole}
		for i, count := range stat {
			fmt.Printf("player %d has %s, won %d times.\n", i, players[i], count)
		}
	*/
}
