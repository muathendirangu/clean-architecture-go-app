package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/muathendirangu/clean-architecture-go-app/usecase"
)

type OrderService interface {
	Items(userID, orderID int) ([]usecase.Item, error)
	Add(userID, orderID, itemID int) error
}

type WebserviceHandler struct {
	OrderService OrderService
}

func (handler WebserviceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
	userID, _ := strconv.Atoi(req.FormValue("userId"))
	orderID, _ := strconv.Atoi(req.FormValue("orderId"))
	items, _ := handler.OrderService.Items(userID, orderID)
	for _, item := range items {
		io.WriteString(res, fmt.Sprintf("item id: %d\n", item.ID))
		io.WriteString(res, fmt.Sprintf("item name: %v\n", item.Name))
		io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
	}
}
