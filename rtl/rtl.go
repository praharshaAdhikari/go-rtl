package rtl

type FixedPoint int64

func (fp *FixedPoint) toFloat(scale int) float64 {
	return float64(*fp) / float64(scale)
}

func (fp *FixedPoint) fromFloat(f float64, scale int) {
	*fp = FixedPoint(f * float64(scale))
}

func (fp *FixedPoint) fromInt(i int, scale int) {
	*fp = FixedPoint(i * int(scale))
}

func (fp *FixedPoint) toInt(scale int) int {
	return int(*fp) / scale
}

type Signal interface {
	Get() FixedPoint
	Set(FixedPoint)
}

type Wire struct {
	value FixedPoint
}

type Register struct {
	value FixedPoint
	next  FixedPoint
}

func NewWire(value ...FixedPoint) *Wire {
	if len(value) > 0 {
		return &Wire{value: value[0]}
	}
	return &Wire{value: 0}
}

func (w *Wire) Get() FixedPoint {
	return w.value
}
func (w *Wire) Set(value FixedPoint) {
	w.value = value
}

func (r *Register) Get() FixedPoint {
	return r.value
}
func (r *Register) Set(value FixedPoint) {
	r.next = value
}
func (r *Register) Clock() {
	r.value = r.next
}
