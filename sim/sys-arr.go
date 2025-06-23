package sim

import (
	"fmt"
	"sync"

	"github.com/praharshaAdhikari/go-rtl/rtl"
)

type SystolicArrayInput [][]MACInput

type SystolicArray struct {
	rows, cols int
	cells      [][]*MAC
	cycles     int
}

func NewSystolicArray(rows, cols int) *SystolicArray {
	cells := make([][]*MAC, rows)
	for i := range cells {
		cells[i] = make([]*MAC, cols)
		for j := range cells[i] {
			cells[i][j] = NewMAC()
		}
	}
	return &SystolicArray{
		rows:  rows,
		cols:  cols,
		cells: cells,
	}
}

func (sa *SystolicArray) Clock() {
	// Step 1: Clock all MAC cells concurrently
	var wg sync.WaitGroup
	for i := range sa.rows {
		for j := range sa.cols {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sa.cells[i][j].Clock()
			}(i, j)
		}
	}
	wg.Wait()

	// Step 2: Capture all current values first to avoid race conditions
	// during propagation
	aValues := make([][]rtl.FixedPoint, sa.rows)
	bValues := make([][]rtl.FixedPoint, sa.rows)

	for i := range sa.rows {
		aValues[i] = make([]rtl.FixedPoint, sa.cols)
		bValues[i] = make([]rtl.FixedPoint, sa.cols)
		for j := range sa.cols {
			aValues[i][j] = sa.cells[i][j].a.Get()
			bValues[i][j] = sa.cells[i][j].b.Get()
		}
	}

	// Step 3: Propagate values using captured values
	for i := sa.rows - 1; i >= 0; i-- {
		for j := sa.cols - 1; j >= 0; j-- {
			if j > 0 {
				sa.cells[i][j].a.Set(aValues[i][j-1])
			}
			if i > 0 {
				sa.cells[i][j].b.Set(bValues[i-1][j])
			}
		}
	}

	sa.cycles++
}

func (sa *SystolicArray) Reset() {
	for i := range sa.rows {
		for j := range sa.cols {
			sa.cells[i][j].a.Set(0)
			sa.cells[i][j].b.Set(0)
			sa.cells[i][j].accumulator.Set(0)
		}
	}
	sa.cycles = 0
}

func (sa *SystolicArray) LoadAccumulators(accumulators [][]rtl.FixedPoint) {
	if len(accumulators) != sa.rows || len(accumulators[0]) != sa.cols {
		panic("Invalid accumulator dimensions")
	}
	for i := range sa.rows {
		for j := range sa.cols {
			sa.cells[i][j].accumulator.Set(accumulators[i][j])
			sa.cells[i][j].accumulator.Clock()
		}
	}
	sa.Clock()
}

func (sa *SystolicArray) Simulate(inputs SystolicArrayInput, verbose bool) {
	if verbose {
		fmt.Println("Starting Systolic Array Simulation")
	}

	if len(inputs) == 0 {
		fmt.Println("No inputs provided for Systolic Array simulation.")
		return
	}

	// Reset all cells (can be done concurrently)
	var wg sync.WaitGroup
	for i := range sa.rows {
		for j := range sa.cols {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sa.cells[i][j].Reset()
			}(i, j)
		}
	}
	wg.Wait()

	if verbose {
		fmt.Println("Initial State:")
		for i := range sa.rows {
			for j := range sa.cols {
				acc := sa.cells[i][j].accumulator.Get()
				fmt.Print("Cell[", i, "][", j, "] A = ", sa.cells[i][j].a.Get(), ", B = ", sa.cells[i][j].b.Get(), ", Accumulator = ", acc, "\n")
			}
		}
	}

	// fmt.Println("Starting Feed Cycles...")
	for inputCycle := range inputs {
		// Set inputs concurrently
		wg.Add(sa.rows + sa.cols) // Maximum number of border cells
		for row := range sa.rows {
			go func(row int) {
				defer wg.Done()
				if row < len(inputs[inputCycle]) && 0 < len(inputs[inputCycle][row]) {
					sa.cells[row][0].a.Set(inputs[inputCycle][row][0][0].Get())
				}
			}(row)
		}

		for col := range sa.cols {
			go func(col int) {
				defer wg.Done()
				if 0 < len(inputs[inputCycle]) && col < len(inputs[inputCycle][0]) {
					sa.cells[0][col].b.Set(inputs[inputCycle][0][col][1].Get())
				}
			}(col)
		}
		wg.Wait()

		// Clock the entire array
		sa.Clock()
		if verbose {
			fmt.Print("Feed Cycle ", inputCycle+1, ":\n")
			for i := range sa.rows {
				for j := range sa.cols {
					acc := sa.cells[i][j].accumulator.Get()
					fmt.Print("Cell[", i, "][", j, "] A = ", sa.cells[i][j].a.Get(), ", B = ", sa.cells[i][j].b.Get(), ", Accumulator = ", acc, "\n")
				}
			}
			fmt.Println("End of Feed Cycle", inputCycle+1)
		}
	}

	// fmt.Println("Starting Drain Cycles...")
	drainCycles := sa.rows + sa.cols - 2 // Correct number of drain cycles
	for drainCycle := range drainCycles {
		sa.Clock()
		if verbose {
			fmt.Printf("Drain Cycle %d:\n", drainCycle+1)
			for i := range sa.rows {
				for j := range sa.cols {
					acc := sa.cells[i][j].accumulator.Get()
					fmt.Print("Cell[", i, "][", j, "] A = ", sa.cells[i][j].a.Get(), ", B = ", sa.cells[i][j].b.Get(), ", Accumulator = ", acc, "\n")
				}
			}
			fmt.Println("End of Drain Cycle", drainCycle+1)
		}
	}

	if verbose {
		fmt.Println("Total Cycles:", sa.cycles)
		fmt.Println("Final MAC States:")
		for i := range sa.rows {
			for j := range sa.cols {
				acc := sa.cells[i][j].accumulator.Get()
				fmt.Print("Cell[", i, "][", j, "] A = ", sa.cells[i][j].a.Get(), ", B = ", sa.cells[i][j].b.Get(), ", Accumulator = ", acc, "\n")
			}
		}
		fmt.Println("Systolic Array Simulation Complete")
	}
}

func (sa *SystolicArray) GetFinalAccumulators() [][]rtl.FixedPoint {
	finalAccumulators := make([][]rtl.FixedPoint, sa.rows)
	for i := range finalAccumulators {
		finalAccumulators[i] = make([]rtl.FixedPoint, sa.cols)
		for j := range finalAccumulators[i] {
			finalAccumulators[i][j] = sa.cells[i][j].accumulator.Get()
		}
	}
	return finalAccumulators
}
