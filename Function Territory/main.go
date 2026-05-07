package main

import "fmt"

/*
 * TEIL 1: Die Basis-Auftriebsfunktion
 * Ein Falke benötigt eine Grundkraft, um in der Luft zu bleiben.
 * Diese berechnet sich aus Windgeschwindigkeit + Anflugwinkel.
 */

// In der Funktionsdefinition fehlen die 2 Datentypen Angaben der Parameter --> ergänzen Sie diese (beides Integer)
func berechneBasisAuftrieb(windgeschwindigkeit int, anflugwinkel int) int {
	// Add values
	ergebnis := windgeschwindigkeit + anflugwinkel

	return ergebnis
}

/*
 * TEIL 2: Der Energieverbrauch
 * Fliegen verbraucht Energie. Je schwerer der Falke, desto mehr.
 * Formel: (Gewicht * Strecke) / 10
 */

// Die Funktion bekommt als Parameter die zwei Integer gewicht & strecke --> passen Sie die Funktionsdefinition entsprechend an
func berechneEnergieVerbrauch(gewicht int, strecke int) int {

	return (gewicht * strecke) / 10
}

/*
 * TEIL 3: Die finale Flug-Validierung (Multiple Return)
 * Ein Falke kann nur fliegen, wenn der Auftrieb höher ist als der Verbrauch.
 * Die Funktion nimmt zwei Datentypen, auftrieb (int) und verbrauch (int) als Parameter
 * Diese Funktion soll zwei Werte zurückgeben:
 * 1. Die Differenz (Netto-Kraft) als Integer
 * 2. Einen Status-Text (string)
 */

// Die Definiton der Paramter, deren Datentypen und die Datentypenangaben der beiden Rückgabewerte fehlen --> ergänzen Sie diese
func checkFlugStatus(auftrieb int, verbrauch int) (int, string) {
	differenz := auftrieb - verbrauch
	status := ""

	if differenz > 0 {
		status = "Stabil"
	} else {
		status = "Absturzgefahr"
	}

	// Return values
	return differenz, status
}

/*
	TEIL 4: Das freudige Kreischen des Falkens
 	Der Falke kann nun wieder fliegen und ist darüber natürlich hoch erfreut!
  	Aufgabe:
   		Erstellen Sie gleich unter diesem Kommentar eine Funktion (ohne Aufruf oder Rückgabeparameter) namens "kreischen"
     	die nichts anderes Tut als ein "KKKRREEEIIIISSSCCCHHHHH" zu printen
      	Rufen Sie diese Funktion in der Main Funktion 3x auf.
*/

func kreischen() {
	fmt.Println("KKKRREEEIIIISSSCCCHHHHH")
}

func main() {
	fmt.Println("--- Analyse der Variable-Falken startet ---")

	// 1. Aufruf der Basis-Funktion
	speed := 45
	winkel := 12
	aktuellerAuftrieb := berechneBasisAuftrieb(speed, winkel)
	fmt.Printf("Berechneter Auftrieb: %d Einheiten\n", aktuellerAuftrieb)

	// 2. Aufruf der Energie-Funktion
	gewicht := 80
	distanz := 20
	verbrauch := berechneEnergieVerbrauch(gewicht, distanz)
	fmt.Printf("Voraussichtlicher Verbrauch: %d Einheiten\n", verbrauch)

	// 3. Finale Prüfung (Auffangen von zwei Rückgabewerten)
	kraftReserve, flugZustand := checkFlugStatus(aktuellerAuftrieb, verbrauch)

	fmt.Println("-------------------------------------------")
	fmt.Printf("Analyse-Ergebnis: %s (Reserve: %d)\n", flugZustand, kraftReserve)

	if flugZustand == "Stabil" && kraftReserve > 0 {
		fmt.Println("Mission abgeschlossen: Der Falke hält seine Bahn!")
	} else {
		fmt.Println("Fehler: Die mathematische Kapselung ist noch instabil.")
	}

	kreischen()
	kreischen()
	kreischen()
}
