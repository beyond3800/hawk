package faker

func Int(min, max int) int {
	return rnd.Intn(max-min+1) + min
}

func Bool() bool {
	return rnd.Intn(2) == 0
}

func Float(min, max float64) float64 {
	return min + rnd.Float64()*(max-min)
}