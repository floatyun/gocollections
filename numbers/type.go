package numbers

type AllInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32
}

type AllFloat interface {
	~float32 | ~float64
}

type AllNumber interface {
	AllInt | AllFloat
}
