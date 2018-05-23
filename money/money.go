package main

type Money struct {
	Amount   int
	Currency string
}

func (m Money) Times(i int) Money {
	return Money{m.Amount * i, m.Currency}
}

func (m Money) Equal(other Money) bool {
	return m.Amount == other.Amount && m.Currency == other.Currency
}
