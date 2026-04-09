package main

import "fmt"

/* STORY-KONTEXT:
   Um den reissenden Fluss ins Function Territory zu überqueren, müssen die Helden
   eine Brücke konstruieren. Sie nutzen IF/ELSE für die Materialprüfung,
   SWITCH für die Pfeilerarten und LOOPS für die tausenden Haltekabel.
*/

// --- TEIL 1: Architektur-Veranschaulichung (Strukturen & Embedding) ---
// Studieren Sie diesen Abschnitt und versuchen Sie zu verstehen, wie das hier "funktioniert!"

type Baumaterial struct {
	Name             string
	Qualitaet        int // Wert von 0 bis 100
}

type BrueckenPfeiler struct {
	Baumaterial           // Embedding: Ein Pfeiler besteht aus Material
	PfeilerTyp   string   // z.B. "Stein", "Holz", "Stahl"
	IstStabil    bool
}

func main() {
	fmt.Println("=== BAU-PROTOKOLL: BRÜCKE DER LOGIK ===")

	// --- TEIL 2: Initialisierung der Objekte ---
	pfeilerKern := BrueckenPfeiler{
		Baumaterial: Baumaterial{
			Name:      "Granit-Block",
			Qualitaet: 85,
		},
		PfeilerTyp: "Stein",
		IstStabil:  false,
	}

	// Variablen für die Bausteuerung
	var materialQualitaet int = pfeilerKern.Qualitaet
	var pfeilerTyp string = pfeilerKern.PfeilerTyp
	var kabelAnzahl int = 5 // Wie viele Haltekabel gespannt werden müssen

	// --- TEIL 3: IHRE IMPLEMENTIERUNG ---

	// AUFGABE A: Materialprüfung (If-Else)
	// Prüfen Sie, ob die 'materialQualitaet' grösser oder gleich 80 ist.

	if materialQualitaet __________ {
		fmt.Println("[CHECK]: Materialprüfung bestanden. Fundament wird gesetzt.")
		pfeilerKern.IstStabil = true

		// AUFGABE B: Pfeiler-Konstruktion (Switch)
		// Nutzen Sie einen Switch für die Variable 'pfeilerTyp'.
		// - Case "Stein": Geben Sie "Massiver Steinpfeiler verankert." aus.
		// - Case "Holz":  Geben Sie "Warnung: Holzpfeiler könnten bei Flut brechen!" aus.
		// - Default:      Geben Sie "Unbekanntes Material! Stabilität nicht garantiert." aus.

		// --- IMPLEMENTIERUNG HIER (Switch-Block) ---


		// AUFGABE C: Haltekabel spannen (For-Loop)
		// Nutzen Sie eine For-Schleife, um die 'kabelAnzahl' nacheinander zu spannen.
		// Geben Sie pro Durchgang aus: "Haltekabel Nr. [X] wird festgezogen..."
                // Schauen Sie dabei, dass X immer die Zahl des aktuellen Durchlaufs angibt. Also 1,2,3 usw..

		fmt.Println("\n[AKTION]: Beginne mit dem Spannen der Kabel...")
		// --- IMPLEMENTIERUNG HIER (For-Loop) ---


	} else {
		// Wenn die Qualität unter 80 liegt, geben Sie aus: "Baugenehmigung verweigert: Material zu schwach!"

	}

	// --- TEIL 4: Abschlussbericht ---
	fmt.Printf("\n--- Konstruktion der Einheit %s abgeschlossen ---\n", pfeilerKern.Name)
}
