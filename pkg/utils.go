package utils

import (
	"sync"

	"github.com/praharshaAdhikari/go-rtl/rtl"
	"github.com/praharshaAdhikari/go-rtl/sim"
)

func MatricesToSystolicArrayInput(matrices [][][]rtl.FixedPoint) sim.SystolicArrayInput {
	A := matrices[0] // MxK
	B := matrices[1] // KxN

	M := len(A)
	K := len(A[0])
	N := len(B[0])

	if len(B) != K {
		panic("Invalid matrices: A's columns must match B's rows")
	}

	cycles := M + K + N - 2 // Correct number of cycles needed
	inputs := make(sim.SystolicArrayInput, cycles)

	var wg sync.WaitGroup

	// Process each cycle concurrently
	for t := range cycles {
		wg.Add(1)
		go func(t int) {
			defer wg.Done()

			row := make([]sim.MACInput, M)
			for i := range M {
				col := make(sim.MACInput, N)
				for j := range N {
					var aSig, bSig rtl.Signal = rtl.NewWire(0), rtl.NewWire(0)

					// Get correct A value for this cycle
					if j == 0 && i <= t && t-i < K {
						aSig = rtl.NewWire(A[i][t-i])
					}

					// Get correct B value for this cycle
					if i == 0 && j <= t && t-j < K {
						bSig = rtl.NewWire(B[t-j][j])
					}

					col[j] = [2]rtl.Signal{aSig, bSig}
				}
				row[i] = col
			}
			inputs[t] = row
		}(t)
	}

	wg.Wait() // Wait for all cycles to be processed
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
