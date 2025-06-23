package main

import (
	utils "github.com/praharshaAdhikari/go-rtl/pkg"
	"github.com/praharshaAdhikari/go-rtl/rtl"
	"github.com/praharshaAdhikari/go-rtl/sim"
)

func main() {
	mac := sim.NewMAC()
	macInputs := sim.MACInput{
		{rtl.NewWire(1), rtl.NewWire(2)},
		{rtl.NewWire(3), rtl.NewWire(4)},
		{rtl.NewWire(5), rtl.NewWire(6)},
		{rtl.NewWire(7), rtl.NewWire(8)},
	}
	mac.Simulate(macInputs, false)

	sysArr := sim.NewSystolicArray(2, 1)
	// accumulators := [][]rtl.FixedPoint{
	// 	{rtl.FixedPoint(10), rtl.FixedPoint(20), rtl.FixedPoint(30)},
	// 	{rtl.FixedPoint(30), rtl.FixedPoint(40), rtl.FixedPoint(50)},
	// }
	matrices := [][][]rtl.FixedPoint{
		{
			{rtl.FixedPoint(1 * rtl.SCALE), rtl.FixedPoint(1 * rtl.SCALE), rtl.FixedPoint(1 * rtl.SCALE)},
			{rtl.FixedPoint(1 * rtl.SCALE), rtl.FixedPoint(1 * rtl.SCALE), rtl.FixedPoint(1 * rtl.SCALE)},
		},
		{
			{rtl.FixedPoint(1 * rtl.SCALE)},
			{rtl.FixedPoint(1 * rtl.SCALE)},
			{rtl.FixedPoint(1 * rtl.SCALE)},
		},
	}
	// sysArr.LoadAccumulators(accumulators)
	sysArrInputs := utils.MatricesToSystolicArrayInput(matrices)
	sysArr.Simulate(sysArrInputs, false)
}
