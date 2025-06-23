package sim

import (
	"fmt"

	"github.com/praharshaAdhikari/go-rtl/rtl"
)

type MACInput [][2]rtl.Signal

type MAC struct {
	a, b        rtl.Signal
	accumulator *rtl.Register
}

func NewMAC() *MAC {
	return &MAC{
		a:           &rtl.Wire{},
		b:           &rtl.Wire{},
		accumulator: &rtl.Register{},
	}
}

func (m *MAC) Clock() {
	product := rtl.Multiply(m.a, m.b)
	result := rtl.Add(m.accumulator, &product)
	m.accumulator.Set(result.Get())
	m.accumulator.Clock()
}

func (m *MAC) Reset() {
	m.a.Set(0)
	m.b.Set(0)
	m.accumulator.Set(0)
}

func (m *MAC) Simulate(inputs MACInput) {
	fmt.Println("Starting MAC Simulation")
	if len(inputs) == 0 {
		fmt.Println("No inputs provided for MAC simulation.")
		return
	}
	fmt.Println("Initial State: A =", m.a.Get(), ", B =", m.b.Get(), ", Accumulator =", m.accumulator.Get())
	for i, input := range inputs {
		m.a.Set(input[0].Get())
		m.b.Set(input[1].Get())
		m.Clock()
		fmt.Print("Cycle ", i+1, ": A = ", m.a.Get(), ", B = ", m.b.Get(), ", Accumulator = ", m.accumulator.Get(), "\n")
	}
	fmt.Println("Final State: A =", m.a.Get(), ", B =", m.b.Get(), ", Accumulator =", m.accumulator.Get())
	fmt.Println("Total Cycles:", len(inputs))
	fmt.Println("MAC Simulation Complete")
}
