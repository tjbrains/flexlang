// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/pkg/flexlang/functions"
)

func TestMathFunctions_Expm1(t *testing.T) {
	var mathFunctions functions.MathFunctions

	t.Log(mathFunctions.Expm1(144.0))
}

func TestMathFunctions_Floor(t *testing.T) {
	var mathFunctions functions.MathFunctions

	assert.Equal(t, int64(1), mathFunctions.Floor(1.0))
	assert.Equal(t, int64(1), mathFunctions.Floor(1.0234))
	assert.Equal(t, int64(1), mathFunctions.Floor(1.5))
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
	assert.Equal(t, int64(34), mathFunctions.Trunc(34.0123))
}
