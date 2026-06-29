package migration

import (
	"fmt"
	"os"
)

func Scan() error{
	var err error
	files, err := os.ReadDir("database/migrations")
	if err != nil{
		return fmt.Errorf("Unable to read Migrations")
	} 
	for _, file := range files{
		content, err := os.ReadFile("database/migrations/"+file.Name())
		fmt.Println(string(content),"scan")
		if err != nil{
			return fmt.Errorf("Unable to read file")
		}
	}
	return nil
}