package faker

import "fmt"

func Phone() string {

	return fmt.Sprintf(
		"080%d%d%d%d",
		Int(100, 999),
		Int(100, 999),
		Int(10, 99),
		Int(10, 99),
	)
}