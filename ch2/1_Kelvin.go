// 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，
// Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。

package main

// Package tempconv performs Celsius and Fahrenheit conversions.
import "fmt"

type Celsius float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	AbsoluteZeroK Kelvin  = 0
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (k Kelvin) String() string  { return fmt.Sprintf("%gK", k) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

func main() {
	fmt.Printf("%v -> %v\n", AbsoluteZeroC, CToK(AbsoluteZeroC))
	fmt.Printf("%v -> %v\n", AbsoluteZeroK, KToC(AbsoluteZeroK))
}
