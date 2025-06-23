package rtl

import "math"

func FixedPointMul(a, b FixedPoint) FixedPoint {
	aFloat := a.ToFloat()
	bFloat := b.ToFloat()
	var result FixedPoint
	result.FromFloat(aFloat * bFloat)
	return result
}

func FixedPointDiv(a, b FixedPoint) FixedPoint {
	if b.ToFloat() == 0 {
		return a
	}
	aFloat := a.ToFloat()
	bFloat := b.ToFloat()
	var result FixedPoint
	result.FromFloat(aFloat / bFloat)
	return result
}

func FixedPointExp(a FixedPoint) FixedPoint {
	aFloat := a.ToFloat()
	expFloat := float32(math.Exp(float64(aFloat)))
	var result FixedPoint
	result.FromFloat(expFloat)
	return result
}
