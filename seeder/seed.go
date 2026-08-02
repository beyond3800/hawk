package seeder



func Run( fn func() error) error{
	return fn()
}