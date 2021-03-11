package slices

func SumFloat64(list []float64) float64 {
	var s float64
	for _, v := range list {
		s = s + v
	}
	return s
}

func SumFloat32(list []float32) float32 {
	var s float32
	for _, v := range list {
		s = s + v
	}
	return s
}
