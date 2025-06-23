package utils

import (
	"github.com/praharshaAdhikari/go-rtl/rtl"
	"github.com/praharshaAdhikari/go-rtl/sim"
)

func MatricesToSystolicArrayInput(matrices [][][]int) sim.SystolicArrayInput {
	A := matrices[0] // MxK
	B := matrices[1] // KxN

	M := len(A)
	K := len(A[0])
	N := len(B[0])

	if len(B) != K {
		panic("Invalid matrices: A's columns must match B's rows")
	}

	cycles := K + N - 1
	inputs := make(sim.SystolicArrayInput, cycles)

	for t := range cycles {
		row := make([]sim.MACInput, M)
		for i := range M {
			col := make(sim.MACInput, N)
			for j := range N {
				var aSig, bSig rtl.Signal = &rtl.Wire{}, &rtl.Wire{}
				aSig.Set(0)
				bSig.Set(0)
				if j == 0 && t-i >= 0 && t-i < K {
					aSig.Set(A[i][t-i])
				}
				if i == 0 && t-j >= 0 && t-j < K {
					bSig.Set(B[t-j][j])
				}
				col[j] = [2]rtl.Signal{aSig, bSig}
			}
			row[i] = col
		}
		inputs[t] = row
	}
	return inputs
}

func PrintSystolicArrayInput(inputs sim.SystolicArrayInput) {
	for cycle, row := range inputs {
		for i, col := range row {
			for j, sigPair := range col {
				a := sigPair[0].Get()
				b := sigPair[1].Get()
				println("Cycle", cycle, "Row", i, "Col", j, "A =", a, "B =", b)
			}
		}
		println("End of Cycle", cycle)
	}
}
