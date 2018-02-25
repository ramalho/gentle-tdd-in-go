package main

type Money struct {
	Ammount int
	Currency string
}

func (v Money) Times(x int) Money {
	return Money{v.Ammount * x, v.Currency}
}
