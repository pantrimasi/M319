package main

import "fmt"

// --- TEIL 1: Architektur-Veranschaulichung (Strukturen & Embedding) ---
// Studieren Sie diesen Abschnitt und versuchen Sie zu verstehen, wie das hier "funktioniert!"

type Baumaterial struct {
	Name      string
	Qualitaet int // Wert von 0 bis 100
}

type BrueckenPfeiler struct {
	Baumaterial        // Embedding: Ein Pfeiler besteht aus Material
	PfeilerTyp  string // z.B. "Stein", "Holz", "Stahl"
	IstStabil   bool
}

func main() {
	fmt.Println("=== BAU-PROTOKOLL: BRÜCKE DER LOGIK ===")

	// --- TEIL 2: Initialisierung der Objekte ---
	pfeilerKern := BrueckenPfeiler{
		Baumaterial: Baumaterial{
			Name:      "Granit-Block",
			Qualitaet: 85,
		},
		PfeilerTyp: "Stahl",
		IstStabil:  true,
	}

	// Variablen für die Bausteuerung
	var materialQualitaet int = pfeilerKern.Qualitaet
	var pfeilerTyp string = pfeilerKern.PfeilerTyp
	var kabelAnzahl int = 5 // Wie viele Haltekabel gespannt werden müssen

	// --- TEIL 3: IHRE IMPLEMENTIERUNG ---
	//
	if materialQualitaet >= 80 {
		fmt.Println("[CHECK]: Materialprüfung bestanden. Fundament wird gesetzt.")
		pfeilerKern.IstStabil = true

		switch pfeilerTyp {
		case "Stein":
			fmt.Println("Massiver Steinpfeiler verankert.")
		case "Holz":
			fmt.Println("Warnung: Holzpfeiler könnten bei Flut brechen!")
		default:
			fmt.Println("Unbekanntes Material! Stabilität nicht garantiert.")
		}

		for i := 1; i <= kabelAnzahl; i++ {
			fmt.Printf("Haltekabel Nr. %d wird festgezogen...\n", i)
		}

	} else {
		// Wenn die Qualität unter 80 liegt, geben Sie aus: "Baugenehmigung verweigert: Material zu schwach!"
		fmt.Println("Baugenehmigung verweigert: Material zu schwach!")
	}

	// --- TEIL 4: Abschlussbericht ---
	fmt.Printf("\n--- Konstruktion der Einheit %s abgeschlossen ---\n", pfeilerKern.Name)
}
