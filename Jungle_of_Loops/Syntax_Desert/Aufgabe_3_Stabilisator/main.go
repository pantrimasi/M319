package main

import "fmt"

func main() {
	fmt.Println("=== Logik-Viadukt von Null-Pointer-City ===")

	// AUFGABE 1: Initialisierung der Umgebungsvariablen
	wasserstand := 75.5
	temperatur := 28
	stadtviertel := "Markt"

	fmt.Printf("\nStatus-Check: %.1f%% Wasser, %d°C Außentemperatur\n", wasserstand, temperatur)
	fmt.Println("-------------------------------------------------")

	// AUFGABE 2: Separate Sicherheits-Checks
	sicherheitGewaehrleistet := true

	if wasserstand > 90 {
		fmt.Println("[WARNUNG]: Wasserstand kritisch! Überlauf-Gefahr.")
		sicherheitGewaehrleistet = false
	}

	if temperatur > 30 {
		fmt.Println("[WARNUNG]: Hitze-Alarm! Hohe Verdunstungsrate.")
		sicherheitGewaehrleistet = false
	}

	// AUFGABE 3: Wasserverteilung nur bei Sicherheit
	if sicherheitGewaehrleistet {
		switch stadtviertel {
		case "Markt":
			fmt.Println("Wasser zum Markt geleitet.")
		case "Wohngebiet":
			fmt.Println("Wasser zum Wohngebiet geleitet.")
		default:
			fmt.Println("Unbekanntes Viertel!")
		}
	} else {
		fmt.Println("System gesperrt!")
	}

	fmt.Println("-------------------------------------------------")
	fmt.Println("=== Viadukt-Steuerung beendet ===")
}
