package sim

import "fmt"

type SystolicArrayInput [][]MACInput

type SystolicArray struct {
	rows, cols int
	cells      [][]*MAC
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
	for i := range sa.rows {
		for j := range sa.cols {
			sa.cells[i][j].Clock()
		}
	}

	for i := sa.rows - 1; i >= 0; i-- {
		for j := sa.cols - 1; j >= 0; j-- {
			if j > 0 {
				sa.cells[i][j].a.Set(sa.cells[i][j-1].a.Get())
			}
			if i > 0 {
				sa.cells[i][j].b.Set(sa.cells[i-1][j].b.Get())
			}
		}
	}
}

func (sa *SystolicArray) Simulate(inputs SystolicArrayInput) {
	for inputCycle := range inputs {
		for row := range sa.rows {
			for col := range sa.cols {
				if row == 0 {
					sa.cells[row][col].b.Set(inputs[inputCycle][row][col][1].Get())
				}
				if col == 0 {
					sa.cells[row][col].a.Set(inputs[inputCycle][row][col][0].Get())
				}
			}
		}

		sa.Clock()
		fmt.Print("Feed Cycle ", inputCycle+1, ":\n")
		for i := range sa.rows {
			for j := range sa.cols {
				acc := sa.cells[i][j].accumulator.Get()
				fmt.Print("Cell[", i, "][", j, "] A = ", sa.cells[i][j].a.Get(), ", B = ", sa.cells[i][j].b.Get(), ", Accumulator = ", acc, "\n")
			}
		}
		fmt.Println("End of Feed Cycle", inputCycle+1)
	}

	drainCycles := sa.rows - 1
	for drainCycle := range drainCycles {
		sa.Clock()
		fmt.Printf("Drain Cycle %d:\n", drainCycle+1)
		for i := range sa.rows {
			for j := range sa.cols {
				acc := sa.cells[i][j].accumulator.Get()
				fmt.Print("Cell[", i, "][", j, "] A = ", sa.cells[i][j].a.Get(), ", B = ", sa.cells[i][j].b.Get(), ", Accumulator = ", acc, "\n")
			}
		}
		fmt.Println("End of Drain Cycle", drainCycle+1)
	}

	fmt.Println("Final MAC States:")
	for i := range sa.rows {
		for j := range sa.cols {
			acc := sa.cells[i][j].accumulator.Get()
			fmt.Print("Cell[", i, "][", j, "] A = ", sa.cells[i][j].a.Get(), ", B = ", sa.cells[i][j].b.Get(), ", Accumulator = ", acc, "\n")
		}
	}
}
