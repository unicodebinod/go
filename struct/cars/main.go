package main

import "fmt"

func main() {
    // Erstellen eines neuen Autos mit den gegebenen Eigenschaften
    car := Cars{
        Marke: "Toyota",
        Model: "Camry",
        Baujahr: 2018,
        PS: 170,
        Farbe: "Silber",
        Verkaufspreis: 20000,
        Sitze: 5,
        Türen: 4,
        Klimaanlage: true,
    }

    // Berechnung des Verbrauchs und Aktualisierung des Autos
    car.SetVerbrauch(11.0, 8.0, 100.0)

    // Ausgabe der Auto-Eigenschaften und des Verbrauchs
    fmt.Printf("Marke: %s\n", car.Marke)
    fmt.Printf("Model: %s\n", car.Model)
    fmt.Printf("Baujahr: %d\n", car.Baujahr)
    fmt.Printf("PS: %d\n", car.PS)
    fmt.Printf("Farbe: %s\n", car.Farbe)
    fmt.Printf("Verkaufspreis: %.2f\n", car.Verkaufspreis)
    fmt.Printf("Sitze: %d\n", car.Sitze)
    fmt.Printf("Türen: %d\n", car.Türen)
    fmt.Printf("Klimaanlage: %t\n", car.Klimaanlage)
    fmt.Printf("Verbrauch: %.2f\n", car.Verbrauch)
}
