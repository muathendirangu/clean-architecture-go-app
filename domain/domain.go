package domain

import "errors"

type CustomerRespository interface {
	Store(customer Customer)
	FindByID(ID int) Customer
}

type ItemRespository interface {
	Store(item Item)
	FindByID(ID int) Item
}

type OrderRespository interface {
	Store(order Order)
	FindByID(id int) Order
}

type Customer struct {
	ID   int
	Name string
}

type Item struct {
	ID        int
	Name      string
	Value     float64
	Available bool
}

type Order struct {
	ID       int
	Customer Customer
	Items    []Item
}

func (order *Order) Add(item Item) error {
	if !item.Available {
		return errors.New("Item not available")
	}
	if order.checkValue()+item.Value > 250 {
		return errors.New("An order cannot exceed the total value of $250")
	}
	order.Items = append(order.Items, item)
	return nil
}

func (order *Order) checkValue() float64 {
	sum := 0.0
	for i := range order.Items {
		sum = sum + order.Items[i].Value
	}
	return sum
}
