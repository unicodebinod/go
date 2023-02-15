package main

type Cars struct {
    Marke        string
    Model        string
    Baujahr      int
    PS           int
    Farbe        string
    Verkaufspreis float64
    Sitze        int
    Türen        int
    Klimaanlage  bool
    Verbrauch    float64
}

func (c *Cars) SetVerbrauch(ortschaft float64, land float64, geschwindigkeit float64) {
    // Berechnung des Verbrauchs pro 100 KM basierend auf den gegebenen Werten
    c.Verbrauch = ((ortschaft/100.0)*11.0 + (land/100.0)*8.0) / (geschwindigkeit/100.0)
}
