package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Feet float64
type Metre float64

func FtToM(ft Feet) Metre {
	return Metre(ft / 3.2808)
}

func MToFt(m Metre) Feet {
	return Feet(m * 3.2808)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%g ft", ft)
}

func (m Metre) String() string {
	return fmt.Sprintf("%g m", m)
}

var feet = flag.Float64("f", -1, "feet")
var metre = flag.Float64("m", -1, "metre")

func main() {
	flag.Parse()
	if *feet == -1 {
		fmt.Printf("Enter feet: ")
		reader := bufio.NewReader(os.Stdin)
		feetStr, _, _ := reader.ReadLine()
		*feet, _ = strconv.ParseFloat(string(feetStr), 64)
	}
	if *metre == -1 {
		fmt.Printf("Enter metre: ")
		reader := bufio.NewReader(os.Stdin)
		metreStr, _, _ := reader.ReadLine()
		*metre, _ = strconv.ParseFloat(string(metreStr), 64)
	}
	fmt.Printf("%v -> %v\n", Feet(*feet), FtToM(Feet(*feet)))
	fmt.Printf("%v -> %v\n", Metre(*metre), MToFt(Metre(*metre)))
}
