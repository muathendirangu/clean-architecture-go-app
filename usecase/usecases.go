package usecase

import (
	"fmt"

	"github.com/muathendirangu/clean-architecture-go-app/domain"
)

type UserRepository interface {
	Store(user User)
	FindByID(ID int) User
}

type User struct {
	ID       int
	IsAdmin  bool
	Customer domain.Customer
}

type Item struct {
	ID    int
	Name  string
	Value float64
}

type Logger interface {
	Log(message string) error
}

type OrderUseCase struct {
	UserRepository  UserRepository
	OrderRepository domain.OrderRespository
	ItemRepository  domain.ItemRespository
	Logger          Logger
}

func (orderUseCase OrderUseCase) Items(userID, orderID int) ([]Item, error) {
	var items []Item
	user := orderUseCase.UserRepository.FindByID(userID)
	order := orderUseCase.OrderRepository.FindByID(orderID)
	if user.Customer.ID != order.Customer.ID {
		message := "User #%i (customer #%i) "
		message += "is not allowed to see items "
		message += "in order #%i (of customer #%i)"
		err := fmt.Errorf(message,
			user.ID,
			user.Customer.ID,
			order.ID,
			order.Customer.ID)
		orderUseCase.Logger.Log(err.Error())
		items = make([]Item, 0)
		return items, err
	}
	items = make([]Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = Item{item.ID, item.Name, item.Value}
	}
	return items, nil
}

func (orderUseCase OrderUseCase) Add(userID, orderID, itemID int) error {
	var message string
	user := orderUseCase.UserRepository.FindByID(userID)
	order := orderUseCase.OrderRepository.FindByID(orderID)
	if user.Customer.ID != order.Customer.ID {
		message = "User #%i (customer #%i) "
		message += "is not allowed to add items "
		message += "to order #%i (of customer #%i)"
		err := fmt.Errorf(message,
			user.ID,
			user.Customer.ID,
			order.ID,
			order.Customer.ID)
		orderUseCase.Logger.Log(err.Error())
		return err
	}
	item := orderUseCase.ItemRepository.FindByID(itemID)
	err := order.Add(item)
	if err != nil {
		message := "user #i (customer #i) "
		message += "is not allowed to add the item to order #i"
		message += "as user #i as the business rule %s was violated"
		err := fmt.Errorf(message,
			item.ID,
			order.ID,
			user.Customer.ID,
			user.ID,
			err.Error())
		orderUseCase.Logger.Log(err.Error())
	}
	orderUseCase.OrderRepository.Store(order)
	orderUseCase.Logger.Log(fmt.Sprintf(
		"User added item '%s' %v to order %v",
		item.Name, item.ID, order.ID))
	return nil
}

type AdminOrderUsecase struct {
	OrderUseCase
}

func (adminUseCase AdminOrderUsecase) Add(userID, orderID, itemID int) error {
	var message string
	user := adminUseCase.UserRepository.FindByID(userID)
	order := adminUseCase.OrderRepository.FindByID(orderID)
	if !user.IsAdmin {
		message = "User #%i (customer #%i) "
		message += "is not allowed to add items "
		message += "to order #%i (of customer #%i)"
		message += "because he is not an adminstrator"
		err := fmt.Errorf(message,
			user.ID,
			user.Customer.ID,
			order.ID,
			order.Customer.ID)
		adminUseCase.Logger.Log(err.Error())
		return err
	}
	item := adminUseCase.ItemRepository.FindByID(itemID)
	err := order.Add(item)
	if err != nil {
		message := "user #i (customer #i) "
		message += "is not allowed to add the item to order #i"
		message += "as user #i as the business rule %s was violated"
		err := fmt.Errorf(message,
			item.ID,
			order.ID,
			user.Customer.ID,
			user.ID,
			err.Error())
		adminUseCase.Logger.Log(err.Error())
	}
	adminUseCase.OrderRepository.Store(order)
	adminUseCase.Logger.Log(fmt.Sprintf(
		"Admin added item '%s' %v to order %v",
		item.Name, item.ID, order.ID))
	return nil
}
