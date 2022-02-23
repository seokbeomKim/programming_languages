package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%g`C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g`F", f) }

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "`C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "`F":
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {
	var temp = CelsiusFlag("temp", 200.0, "temperature")
	flag.Parse()
	fmt.Println(*temp)
}
