package units

import (
	"fmt"
	"sort"
)

// ---------- Constants ----------
const (
	MetreToFoot   = 3.28084
	FootToMetre   = 0.3048
	KgToLbs       = 2.20462
	LbsToKg       = 0.453592
	LitreToGallon = 0.264172
	GallonToLitre = 3.78541
)

// ---------- Raw converter functions ----------
func metresToFeet(m float64) float64       { return m * MetreToFoot }
func feetToMetres(ft float64) float64     { return ft * FootToMetre }
func celsiusToFahrenheit(c float64) float64 { return c*1.8 + 32 }
func fahrenheitToCelsius(f float64) float64 { return (f - 32) / 1.8 }
func kgToLbs(kg float64) float64          { return kg * KgToLbs }
func lbsToKg(lbs float64) float64        { return lbs * LbsToKg }
func litresToGallons(l float64) float64   { return l * LitreToGallon }
func gallonsToLitres(g float64) float64   { return g * GallonToLitre }

// ---------- Currency exchange rates (stretch goal) ----------
var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.92,
	"GBP": 0.79,
	"JPY": 150.0,
	"CAD": 1.35,
}

// ---------- Conversions registry ----------
// Conversions maps "from->to" strings to closures.
var Conversions = map[string]func(float64) float64{
	"m->ft":  metresToFeet,
	"ft->m":  feetToMetres,
	"c->f":   celsiusToFahrenheit,
	"f->c":   fahrenheitToCelsius,
	"kg->lb": kgToLbs,
	"lb->kg": lbsToKg,
	"l->gal": litresToGallons,
	"gal->l": gallonsToLitres,
}

func init() {
	// Currency conversions as closures (stretch goal)
	// We capture rateFrom and rateTo locally to avoid the loop-variable closure trap.
	for from, rateFrom := range exchangeRates {
		for to, rateTo := range exchangeRates {
			if from == to {
				continue
			}
			key := fmt.Sprintf("%s->%s", from, to)
			rFrom := rateFrom
			rTo := rateTo
			Conversions[key] = func(v float64) float64 {
				return v * (rTo / rFrom)
			}
		}
	}
}

// Convert looks up a conversion and returns an error for unknown pairs.
func Convert(from, to string, value float64) (float64, error) {
	key := fmt.Sprintf("%s->%s", from, to)
	fn, ok := Conversions[key]
	if !ok {
		return 0, fmt.Errorf("unsupported conversion: %s to %s", from, to)
	}
	return fn(value), nil
}

// ListConversions returns all supported conversion keys, sorted.
func ListConversions() []string {
	keys := make([]string, 0, len(Conversions))
	for k := range Conversions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}