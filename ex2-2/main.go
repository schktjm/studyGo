package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Celsius float64
type Fahrenheit float64
type Meters float64
type Feet float64
type Pound float64
type Kilogram float64

var leng = flag.Bool("l", false, "parse length")
var wgt = flag.Bool("w", false, "parse to weight")
var tem = flag.Bool("t", false, "parse to temperture")

func main() {
	if len(os.Args) > 1 {
		flag.Parse()
		val, _ := strconv.ParseFloat(flag.Args()[0], 64)
		switch {
		case *leng:
			printParse("length", val)
		case *wgt:
			printParse("weight", val)
		case *tem:
			printParse("temperture", val)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		fmt.Println("input the param and number(ex: length 4)")
		input.Scan()
		cmd := strings.Split(input.Text(), " ")
		val, _ := strconv.ParseFloat(cmd[1], 64)
		printParse(cmd[0], val)
	}
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func MToF(m Meters) Feet        { return Feet(m * 3.2808) }
func FToM(f Feet) Meters        { return Meters(f / 3.2808) }
func PToKG(p Pound) Kilogram    { return Kilogram(p / 2.205) }
func KGToP(kg Kilogram) Pound   { return Pound(kg * 2.205) }

func printParse(cmd string, val float64) {
	switch {
	case cmd == "length":
		mer := Meters(val)
		fmt.Printf("%.3fm = %.3fft\n", mer, MToF(mer))
		ft := Feet(val)
		fmt.Printf("%.3fft = %.3fm\n", ft, FToM(ft))
	case cmd == "weight":
		p := Pound(val)
		fmt.Printf("%.3flb = %.3fkg\n", p, PToKG(p))
		kg := Kilogram(val)
		fmt.Printf("%.3fkg = %.3flb\n", kg, KGToP(kg))
	case cmd == "temperture":
		cel := Celsius(val)
		fmt.Printf("%.3fC = %.3fF\n", cel, CToF(cel))
		f := Fahrenheit(val)
		fmt.Printf("%.3fF = %.3fC\n", f, FToC(f))
	default:
		fmt.Println("input the param and number(ex: length 4)")
	}
}
