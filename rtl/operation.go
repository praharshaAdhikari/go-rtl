package rtl

func Add(a, b Signal) Wire {
	result := NewWire()
	result.Set(a.Get() + b.Get())
	return *result
}

func Multiply(a, b Signal) Wire {
	result := NewWire()
	result.Set(a.Get() * b.Get())
	return *result
}
