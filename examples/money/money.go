package main

type Money struct {
	amount   int
	currency string
}

func (m Money) Times(i int) Money {
	return Money{m.amount * i, m.currency}
}

func (m Money) Equal(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}
