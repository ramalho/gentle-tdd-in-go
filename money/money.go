package main

type Dollar struct {
	Ammount int
}

func (v *Dollar) Times(x int)  {
	v.Ammount *= x
}
