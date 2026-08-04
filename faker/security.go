package faker

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Password() string {

	length := 12

	b := make([]byte, length)

	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}