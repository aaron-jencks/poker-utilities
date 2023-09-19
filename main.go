package main

import (
	"fmt"

	"github.com/aaron-jencks/poker/statistics"
)

func main() {
	var pn, tot, ss int
	fmt.Printf("Please enter the total number of players: ")
	fmt.Scan(&pn)
	tot = statistics.PossibleHandCount(pn)
	ss = statistics.Slovin(tot, 0.01)
	fmt.Printf("Sample size required for %d possible hands: %d\n", tot, ss)
}
