package main

import "fmt"

func main() {

	var fliesstWasserAmBoden bool = false
	var hatTotenBaumLinks bool = false
	var liegtImSchatten bool = true
	var istBreitUndGerade bool = true

	pruefeCanyon(fliesstWasserAmBoden, hatTotenBaumLinks, liegtImSchatten, istBreitUndGerade)

}

func pruefeCanyon(fliesstWasserAmBoden, hatTotenBaumLinks, liegtImSchatten, istBreitUndGerade bool) {
	if fliesstWasserAmBoden {
		if hatTotenBaumLinks {
			fmt.Println("Ergebnis: Dieser Canyon ist RICHTIG (Führt zum Ziel)")
		} else {
			fmt.Println("Ergebnis: Dieser Canyon ist FALSCH (Führt NICHT zum Ziel)")
		}
	} else {
		if liegtImSchatten {
			if istBreitUndGerade {
				fmt.Println("Ergebnis: Dieser Canyon ist RICHTIG (Führt zum Ziel)")
			} else {
				fmt.Println("Ergebnis: Dieser Canyon ist FALSCH (Führt NICHT zum Ziel)")
			}
		} else {
			fmt.Println("Ergebnis: Dieser Canyon ist FALSCH (Führt NICHT zum Ziel)")
		}
	}
}
