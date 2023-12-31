package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = -17.78
	BoilingF      Fahrenheit = 212
	AbsoluteZeroK Kelvin     = 0
	FreezingK     Kelvin     = 273.15
	BoilingK      Kelvin     = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g*C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g*F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g*K", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
