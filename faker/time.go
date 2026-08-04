package faker

import "time"

func Date() time.Time {

	start := time.Now().AddDate(-5, 0, 0).Unix()
	end := time.Now().Unix()

	return time.Unix(rnd.Int63n(end-start)+start, 0)
}