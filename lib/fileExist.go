package lib

import "os"

func FileExist(dir string, fileName string) bool {
	file := dir+"/"+fileName+".go"
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return false
}