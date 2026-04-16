// definiert den Code-Einstiegspunkt für die Barken-Beschwörung
package main

// importiert das Formatierungs-Paket zur Ausgabe von Text
import "fmt"

// =================================================================
// ARCHITEKTUR-BEREICH (Fortgeschrittene Magie - Bitte nicht ändern)
// =================================================================

// Fahrzeug ist die Basisklasse (Struktur) für alle Fortbewegungsmittel
type Fahrzeug struct {
	Name string
}

// Barke erbt Eigenschaften vom Fahrzeug und erweitert diese
type Barke struct {
	Fahrzeug
	Tragkraft  float64
	Passagiere int
}

// Methode zum Materialisieren der Barke
func (b Barke) beschwoere() {
	fmt.Printf("Die Barke '%s' beginnt zu schimmern...\n", b.Name)
	// Loop zur Visualisierung der Stabilisierung mit dreimaligem Output
	for i := 0; i < 3; i++ {
		fmt.Println("... [DATA-STREAM-STABILIZING] ...")
	}
}

// =================================================================
// IHRE AUFGABE: DIE LOGIK-WERKSTATT
// =================================================================

// Die Haupt-Funktion orchestriert die Barken-Beschwörung und prüft alle Parameter
func main() {
	fmt.Println("=== Die Beschwörung der Logik-Barke ===")

	// AUFGABE 1: Variablen-Initialisierung
	// Deklarieren Sie die Variable 'tragkraft' als float64 mit dem Wert 500.5
	tragkraft := 500.5
	// Deklarieren Sie die Variable 'anzahlPassagiere' als int mit dem Wert 5
	anzahlPassagiere := 5

	// Float64 speichert Dezimalwerte für präzise Gewichtsberechnungen

	/* FIXME: Der nachfolgende Block berechnet die Stabilität.
	   Die Logik der alten Meister ist korrupt.
	   Bedingung: Wenn die anzahlPassagiere größer als 6 ODER
	   die tragkraft kleiner als 400 ist, bricht die Beschwörung ab.
	*/

	istStabil := false

	// Logische Bedingung prüft Instabilität bei zu vielen Passagieren oder niedriger Tragkraft
	if anzahlPassagiere > 6 || tragkraft < 400 {
		fmt.Println("[WARNUNG]: Die Logik-Parameter sind instabil!")
		istStabil = false
	} else {
		fmt.Println("[CHECK]: Auftriebs-Logik stabil.")
		istStabil = true
	}

	// Trennlinie zur besseren Lesbarkeit der Ausgabe
	fmt.Println("---------------------------------------")

	// AUFGABE 2: Navigationstyp festlegen
	// Nutzen Sie einen Switch, um je nach 'anzahlPassagiere' einen Modus auszugeben:
	// 1-3 Passagiere: "Modus: Kleines Boot"
	// 4-6 Passagiere: "Modus: Gruppen-Barke"
	// Default: "Modus: Unbekannt"

	switch anzahlPassagiere {
	case 1, 2, 3:
		fmt.Println("Modus: Kleines Boot")
	case 4, 5, 6:
		fmt.Println("Modus: Gruppen-Barke")
	default:
		fmt.Println("Modus: Unbekannt")
	}

	/* Das Switch-Statement evaluiert die Anzahl der Passagiere und ordnet einen
	   entsprechenden Navigations-Modus zu. Bei 1-3 Personen wird das Boot als klein
	   klassifiziert, bei 4-6 als Gruppenbarke. Alle anderen Werte triggern den
	   Standard-Modus Unbekannt zur Fehlerbehandlung.
	*/

	// FINALER SCHRITT: Wenn alles stabil ist, wird die Barke erzeugt
	if istStabil {
		// Hier wird ein neues Objekt der Klasse 'Barke' erstellt
		meineBarke := Barke{
			Fahrzeug:   Fahrzeug{Name: "Lumen Logica"},
			Tragkraft:  tragkraft,
			Passagiere: anzahlPassagiere,
		}

		meineBarke.beschwoere()
		fmt.Println("Erfolg: Die Barke ist bereit für das Bytestromdelta!")
	} else {
		fmt.Println("Fehler: Die Beschwörung ist fehlgeschlagen.")
	}
}
