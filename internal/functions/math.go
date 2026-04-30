// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"math"
	"math/bits"
	"math/rand/v2"
)

const (
	MathE       = math.E
	MathLN10    = math.Ln10
	MathLOG10E  = math.Log10E
	MathPI      = math.Pi
	MathSQRT1_2 = 1 / math.Sqrt2
	MathSQRT2   = math.Sqrt2
)

// 已忽略的函数：
//  - Math.acos ( x )
//  - Math.acosh ( x )
//  - Math.asin ( x )
//  - Math.asinh ( x )
//  - Math.atan ( x)
//  - Math.asinh ( x )
//  - Math.atan2 ( y, x )
//  - Math.fround ( x )
//  - Math.f16round ( x )
//  - Math.imul ( x, y )
//  - Math.max ( ...args )
//  - Math.min ( ...args )

type MathFunctions struct{}

func (this MathFunctions) Abs(x float64) float64 {
	return math.Abs(x)
}

func (this MathFunctions) Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

func (this MathFunctions) Ceil(x float64) int64 {
	return int64(math.Ceil(x))
}

func (this MathFunctions) Clz32(x uint32) int {
	return bits.LeadingZeros32(x)
}

func (this MathFunctions) Cos(x float64) float64 {
	return math.Cos(x)
}

func (this MathFunctions) Cosh(x float64) float64 {
	return math.Cosh(x)
}

func (this MathFunctions) Exp(x float64) float64 {
	return math.Exp(x)
}

func (this MathFunctions) Expm1(x float64) float64 {
	return math.Expm1(x)
}

func (this MathFunctions) Floor(x float64) int64 {
	return int64(math.Floor(x))
}

func (this MathFunctions) Hypot(args ...float64) float64 {
	var sum float64
	for _, arg := range args {
		sum += math.Pow(arg, 2)
	}

	return math.Sqrt(sum)
}

func (this MathFunctions) Log(x float64) float64 {
	return math.Log(x)
}

func (this MathFunctions) Log1p(x float64) float64 {
	return math.Log1p(x)
}

func (this MathFunctions) Log10(x float64) float64 {
	return math.Log10(x)
}

func (this MathFunctions) Log2(x float64) float64 {
	return math.Log2(x)
}

func (this MathFunctions) Pow(base float64, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func (this MathFunctions) Random() float64 {
	var f = rand.Float64()
	return f - math.Floor(f)
}

func (this MathFunctions) RandN(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.IntN(n)
}

func (this MathFunctions) Round(x float64) float64 {
	return math.Round(x)
}

func (this MathFunctions) Sign(x float64) float64 {
	if math.IsNaN(x) {
		return x
	}

	if x > 0 {
		return 1
	}

	if x < 0 {
		return -1
	}

	return 0
}

func (this MathFunctions) Sin(x float64) float64 {
	return math.Sin(x)
}

func (this MathFunctions) Sinh(x float64) float64 {
	return math.Sinh(x)
}

func (this MathFunctions) Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func (this MathFunctions) Tan(x float64) float64 {
	return math.Tan(x)
}

func (this MathFunctions) Tanh(x float64) float64 {
	return math.Tanh(x)
}

func (this MathFunctions) Trunc(x float64) int64 {
	return int64(math.Trunc(x))
}
