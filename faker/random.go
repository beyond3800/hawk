package faker

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {

	if length <= 0 {
		return ""
	}

	b := make([]byte, length)

	for i := range b {
		b[i] = letters[rnd.Intn(len(letters))]
	}

	return string(b)
}