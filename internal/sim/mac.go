package sim

import (
	"fmt"

	"github.com/praharshaAdhikari/go-rtl/internal/rtl"
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

func (m *MAC) Simulate(inputs MACInput) {
	for i, input := range inputs {
		m.a.Set(input[0].Get())
		m.b.Set(input[1].Get())
		m.Clock()
		fmt.Print("Cycle ", i+1, ": A = ", m.a.Get(), ", B = ", m.b.Get(), ", Accumulator = ", m.accumulator.Get(), "\n")
	}

}
