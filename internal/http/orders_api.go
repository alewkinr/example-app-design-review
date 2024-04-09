package http

import (
	"encoding/json"
	"net/http"

	"github.com/alewkinr/example-app-design-review/internal/orders"
)

type OrdersAPI struct {
	orders *orders.Manager
}

func NewOrdersAPI(orders *orders.Manager) *OrdersAPI {
	return &OrdersAPI{orders: orders}
}

func (api *OrdersAPI) Routes() Routes {
	return Routes{
		"CreateOrder": Route{
			"POST",
			"/orders",
			api.CreateOrderV1,
		},
	}
}

// CreateOrderV1 — POST v1/orders
func (api *OrdersAPI) CreateOrderV1(w http.ResponseWriter, r *http.Request) {
	createOrdReq := &CreateOrderV1Request{}

	if err := json.NewDecoder(r.Body).Decode(createOrdReq); err != nil {
		EncodeJSONResponse(map[string]string{"error": err.Error()}, http.StatusBadRequest, w)
		return
	}

	if validateErr := createOrdReq.Validate(); validateErr != nil {
		EncodeJSONResponse(map[string]string{"error": validateErr.Error()}, http.StatusBadRequest, w)
		return
	}

	updatedOrder, createOrdErr := api.orders.CreateOrder(orders.Order{
		UserEmail:        createOrdReq.UserEmail,
		HotelID:          createOrdReq.HotelID,
		RoomID:           createOrdReq.RoomID,
		CheckInDateTime:  createOrdReq.From,
		CheckOutDateTime: createOrdReq.To,
	})
	if createOrdErr != nil {
		EncodeJSONResponse(createOrdErr.Error(), http.StatusInternalServerError, w)
		return
	}

	if updatedOrder.IsEmpty() {
		EncodeJSONResponse(map[string]string{"error": "room is not available"}, http.StatusForbidden, w)
		return
	}

	// 201 OK
	EncodeJSONResponse(&CreateOrderV1Response{
		HotelID:   updatedOrder.HotelID,
		RoomID:    updatedOrder.RoomID,
		Status:    updatedOrder.Status,
		From:      updatedOrder.CheckInDateTime,
		To:        updatedOrder.CheckOutDateTime,
		UserEmail: updatedOrder.UserEmail,
	}, http.StatusCreated, w)
}
