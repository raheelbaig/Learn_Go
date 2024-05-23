package main

type bill struct {
	name   string
	date   string
	amount float32
}

func NewBill(name string) bill {
	customerBill := bill{
		name:   name,
		date:   "2-5-2024",
		amount: 64000,
	}

	return customerBill
}
