package rtl

type Signal interface {
	Get() int
	Set(int)
}

type Wire struct {
	value int
}

type Register struct {
	value int
	next  int
}

func NewWire(value ...int) *Wire {
	if len(value) > 0 {
		return &Wire{value: value[0]}
	}
	return &Wire{value: 0}
}

func (w *Wire) Get() int {
	return w.value
}
func (w *Wire) Set(value int) {
	w.value = value
}

func (r *Register) Get() int {
	return r.value
}
func (r *Register) Set(value int) {
	r.next = value
}
func (r *Register) Clock() {
	r.value = r.next
}
