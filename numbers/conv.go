package types

// ConvInt T是目标类型，F是来源类型
func ConvInt[T AllInt, F AllInt](f F) T {
	return T(f)
}
