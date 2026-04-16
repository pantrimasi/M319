package main

import "fmt"

// Basis-Struktur für alle Tempel-Elemente
type TempelElement struct {
	Name       string
	Material   string
	MagieFokus bool
}

// Spezialisierung: Eine Wand mit Inschriften
type InschriftWand struct {
	TempelElement // Einbetten der Basis-Eigenschaften
	Sprache       string
	IstEntziffert bool
}

// Eine bereits funktionierende Schutz-Funktion der Wand
func checkWandStatus(wand InschriftWand) {
	if wand.MagieFokus && wand.IstEntziffert {
		fmt.Printf("[INFO]: Die Wand '%s' leuchtet in %s.\n", wand.Name, wand.Sprache)
	}
}

type EnergieKristall struct {
	Farbe        string
	LadungsStand int // 0 bis 100
}

type TempelKern struct {
	EnergieKristall // Embedding: Der Kern enthält den Kristall
	Frequenz        float64
	IsReaktiviert   bool
}

func main() {
	// Vorhandene Wand-Elemente initialisieren (Teil der Tempelwand)
	nordWand := InschriftWand{
		TempelElement: TempelElement{Name: "Wand des Wissens", Material: "Sandstein", MagieFokus: true},
		Sprache:       "Go-Runen",
		IstEntziffert: true,
	}
	checkWandStatus(nordWand)

	fmt.Println("\n=== PROTOKOLL: REAKTIVIERUNG DES TEMPEL-KRISTALLS ===")

	// --- INITIALISIERUNG DES TEMPEL-KERNS ---
	reaktor := TempelKern{
		EnergieKristall: EnergieKristall{
			Farbe:        "Smaragdgrün",
			LadungsStand: 0,
		},
		Frequenz:      440.0,
		IsReaktiviert: false,
	}

	// Variablen für die Aktivierungssequenz
	var frequenz float64 = reaktor.Frequenz
	var aktivierungsCode string = "DRAGON_WEAKEN" // Modi: "DRAGON_WEAKEN", "SAFE_MODE", "SHUTDOWN"
	var zielLadung int = 100

	// --- HIER IHRE IMPLEMENTIERUNG ---

	// AUFGABE A: Frequenz-Check (If-Else)
	// Der Kristall reagiert nur, wenn die Frequenz exakt 440.0 ist
	// UND der LadungsStand kleiner als 100 ist.
	if frequenz _____ && reaktor.LadungsStand _____ {
		fmt.Println("[STATUS]: Frequenz harmonisiert. Initialisiere Ladevorgang..."
		// AUFGABE B: Modus-Wahl für den Kristall (Switch)
		// Nutzen Sie einen Switch für 'aktivierungsCode'.
		// Ergänzen Sie um den zusätzlichen den Case "SHUTDOWN"
		switch aktivierungsCode {
		case "DRAGON_WEAKEN":
			// --- IMPLEMENTIERUNG HIER ---
		case "SAFE_MODE":
			// --- IMPLEMENTIERUNG HIER ---
		default:
			// --- IMPLEMENTIERUNG HIER ---
		}

		// AUFGABE C: Die Lade-Schleife (For-Loop)
		// Erhöhen Sie den 'LadungsStand' von seinem aktuellen Wert bis zum 'zielLadung'
		// in 20er-Schritten (i += 20).
		fmt.Println("\n[AKTION]: Starte Energie-Injektion...")

		for i := reaktor.LadungsStand; i <= _______; i += _____ {
			reaktor.LadungsStand = i
			fmt.Printf("Kristall-Ladung bei %d%%...\n", reaktor.LadungsStand)
		}

		if reaktor.LadungsStand >= 100 {
			reaktor.IsReaktiviert = true
			fmt.Println("\nWWWWUUUUSSSHHHH! Die grüne Lichtsäule schiesst in den Himmel!")
		}

	} else {
		fmt.Println("[FEHLER]: Kristall bleibt dunkel. Überprüfen Sie Frequenz und Energie!")
	}

	// --- ABSCHLUSSBERICHT ---
	fmt.Printf("\n--- Status Bericht für Kristall-Farbe: %s ---", reaktor.Farbe)
	fmt.Printf("\nReaktivierung erfolgreich: %t\n", reaktor.IsReaktiviert)
}
