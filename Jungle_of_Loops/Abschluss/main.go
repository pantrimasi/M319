package main

import "fmt"

func main() {
	// START: Die Helden erreichen das Ufer
	fmt.Println("Die Helden erreichen die Grenze zum Function Territory.")

	// --- ENTSCHEIDUNG (if/else) ---
	// Prüfung, ob die Mana-Quelle rein ist
	istManaRein := true

	if istManaRein {
		fmt.Println("Die Quelle ist rein. Energie fließt in das Siegel.")
	} else {
		fmt.Println("Die Quelle ist getrübt. Die Helden müssen sie reinigen.")
		fmt.Println("Die Quelle wird gereinigt.....")
		fmt.Println("Quelle gereinigt!")
	}

	// --- MEHRFACHENTSCHEIDUNG (switch) ---
	// Auswahl des Segeltyps bestimmt den weiteren Pfad
	segelTyp := "Licht"

	switch segelTyp {
	case "Licht":
		fmt.Println("Segel aus purem Licht werden gewebt.")
	case "Code":
		fmt.Println("Segel aus festen Algorithmen werden konstruiert.")
	default:
		fmt.Println("Einfache Stoffsegel werden gesetzt.")
	}

	// --- SCHLEIFE (for-loop) ---
	// In Go gibt es nur 'for', das hier wie eine 'while'-Schleife fungiert
	energieStatus := 0
	fmt.Println("Beginne mit der Energetisierung des Rumpfes...")

	for energieStatus < 100 {
		fmt.Printf("Energie bei %d%%...\n", energieStatus)
		energieStatus += 25
		// Im Diagramm: Pfeil zurück zur Bedingungsprüfung (Status < 100)
	}

	// ABSCHLUSS
	fmt.Println("Die Galeone ist manifestiert! Die Überfahrt beginnt.")
}
