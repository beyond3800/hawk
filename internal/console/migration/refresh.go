package migration


func Refresh() error{
	err := Reset()
	if err != nil{
		return err
	}
	return Run()
}