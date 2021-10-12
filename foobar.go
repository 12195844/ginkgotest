package bddtest
type Foo struct {
	x int
}

func (f *Foo)Increase() int {
	f.x++
	return f.x
}
func (f *Foo) Add(i int) int  {
	f.x+=i
	return f.x
}
