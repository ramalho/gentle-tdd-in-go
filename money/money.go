package main

type Money struct {
	Ammount int
}

func (v *Money) Times(x int)  {
	v.Ammount *= x
}
