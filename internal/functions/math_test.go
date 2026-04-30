// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/internal/functions"
)

func TestMathFunctions_Abs(t *testing.T) {
	var mathFunctions functions.MathFunctions
	assert.Equal(t, 2.0, mathFunctions.Abs(2.0))
	assert.Equal(t, 2.123, mathFunctions.Abs(-2.123))
}

func TestMathFunctions_Cbrt(t *testing.T) {
	var mathFunctions functions.MathFunctions
	assert.Equal(t, 2.0, mathFunctions.Cbrt(8.0))
	assert.Equal(t, -3.0, mathFunctions.Cbrt(-27.0))
}

func TestMathFunctions_Ceil(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, int64(1), mathFunctions.Ceil(1.0))
	assert.Equal(t, int64(2), mathFunctions.Ceil(1.0234))
	assert.Equal(t, int64(2), mathFunctions.Ceil(1.5))
	assert.Equal(t, int64(124), mathFunctions.Ceil(123.678))
}

func TestMathFunctions_Clz32(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, 31, mathFunctions.Clz32(1))
	assert.Equal(t, 29, mathFunctions.Clz32(4))
	assert.Equal(t, 32, mathFunctions.Clz32(0))
	assert.Equal(t, 0, mathFunctions.Clz32(2147483648))
}

func TestMathFunctions_Cos(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, 0.49999999999999994, mathFunctions.Cos(math.Pi/3))

	t.Log(mathFunctions.Cosh(30))
	t.Log(mathFunctions.Cosh(60))
}

func TestMathFunctions_Exp(t *testing.T) {
	var mathFunctions functions.MathFunctions

	t.Log(mathFunctions.Exp(144.0))
	t.Log(mathFunctions.Exp(100.0))
}

func TestMathFunctions_Expm1(t *testing.T) {
	var mathFunctions functions.MathFunctions

	t.Log(mathFunctions.Expm1(144.0))
	t.Log(mathFunctions.Expm1(100.0))
}

func TestMathFunctions_Floor(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, int64(1), mathFunctions.Floor(1.0))
	assert.Equal(t, int64(1), mathFunctions.Floor(1.0234))
	assert.Equal(t, int64(1), mathFunctions.Floor(1.5))
	assert.Equal(t, int64(123), mathFunctions.Floor(123.678))
}

func TestMathFunctions_Hypot(t *testing.T) {
	var mathFunctions functions.MathFunctions
	assert.Equal(t, 5.0, mathFunctions.Hypot(3, 4))
	assert.Equal(t, 5.0, mathFunctions.Hypot(-3, -4))
	assert.Equal(t, 0.0, mathFunctions.Hypot(0, 0))
}

func TestMathFunctions_Log(t *testing.T) {
	var mathFunctions functions.MathFunctions

	t.Log(mathFunctions.Log(2))
	t.Log(mathFunctions.Log1p(2))
	t.Log(mathFunctions.Log10(2))
	t.Log(mathFunctions.Log2(2))
	t.Log(mathFunctions.Log2(10))
}

func TestMathFunctions_Pow(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, 1.0, mathFunctions.Pow(2.0, 0.0))
	assert.Equal(t, 2.0, mathFunctions.Pow(4.0, 0.5))
	assert.True(t, math.IsNaN(mathFunctions.Pow(-4.0, 0.5)))
}

func TestMathFunctions_Random(t *testing.T) {
	var mathFunctions functions.MathFunctions

	t.Log(mathFunctions.Random())
	assert.True(t, mathFunctions.Random() < 1)
	assert.True(t, mathFunctions.Random() >= 0)

	for range 10_000_000 {
		if mathFunctions.Random() == 0 {
			t.Log("=0")
		}
		if mathFunctions.Random() >= 1 {
			t.Fatal("should not be >= 1")
		}
	}
}

func TestMathFunctions_RandN(t *testing.T) {
	var mathFunctions functions.MathFunctions

	t.Log(mathFunctions.RandN(10))
	assert.True(t, mathFunctions.Random() < 10)
	assert.True(t, mathFunctions.Random() >= 0)

	t.Log(mathFunctions.RandN(0))
	t.Log(mathFunctions.RandN(-10))
}

func TestMathFunctions_Round(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, 2.0, mathFunctions.Round(2.0))
	assert.Equal(t, 2.0, mathFunctions.Round(2.0123))
	assert.Equal(t, 3.0, mathFunctions.Round(2.6123))
}

func TestMathFunctions_Sign(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, 1.0, mathFunctions.Sign(1.0))
	assert.Equal(t, -1.0, mathFunctions.Sign(-123456.123))
	assert.Equal(t, 0.0, mathFunctions.Sign(0))
}

func TestMathFunctions_Sin(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, 0.5, mathFunctions.Sin(math.Pi/6))

	t.Log(mathFunctions.Sinh(30))
	t.Log(mathFunctions.Sinh(60))
}

func TestMathFunctions_Sqrt(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, 2.0, mathFunctions.Sqrt(4))
	assert.Equal(t, 3.0, mathFunctions.Sqrt(9))
}

func TestMathFunctions_Tanh(t *testing.T) {
	var mathFunctions functions.MathFunctions

	t.Log(mathFunctions.Tan(math.Pi / 4))
	t.Log(mathFunctions.Tanh(45))
}

func TestMathFunctions_Trunc(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, int64(2), mathFunctions.Trunc(2.0))
	assert.Equal(t, int64(2), mathFunctions.Trunc(2.0123))
	assert.Equal(t, int64(34), mathFunctions.Trunc(34.6789))
}
