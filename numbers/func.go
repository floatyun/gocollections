package numbers

func Min[V AllNumber](a, b V) V {
	if a < b {
		return a
	}
	return b
}

func Max[V AllNumber](a, b V) V {
	if a > b {
		return a
	}
	return b
}
