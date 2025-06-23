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
	mac.Simulate(macInputs)

	sysArr := sim.NewSystolicArray(2, 3)
	matrices := [][][]rtl.FixedPoint{
		{
			{1, 2},
			{3, 4},
		},
		{
			{9, 8, 7},
			{6, 5, 4},
		},
	}
	sysArrInputs := utils.MatricesToSystolicArrayInput(matrices)
	sysArr.Simulate(sysArrInputs)
}
