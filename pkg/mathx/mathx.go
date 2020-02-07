package mathx

func MaxI(values ...int) (max int) {
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return
}

func MinI(values ...int) (min int) {
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return
}
