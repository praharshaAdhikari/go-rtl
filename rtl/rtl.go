package rtl

type FixedPoint int64

func (fp *FixedPoint) ToFloat(scale int) float32 {
	return float32(*fp) / float32(scale)
}

func (fp *FixedPoint) FromFloat(f float32, scale int) {
	*fp = FixedPoint(f * float32(scale))
}

func (fp *FixedPoint) FromInt(i int, scale int) {
	*fp = FixedPoint(i * int(scale))
}

func (fp *FixedPoint) ToInt(scale int) int {
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
