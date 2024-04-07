package http

import (
	"errors"
	"time"
)

// CreateOrderV1Request — POST v1/orders request model
type CreateOrderV1Request struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

// Validate — basic request validation
func (req *CreateOrderV1Request) Validate() error {
	if req.From.UTC().After(req.To.UTC()) || req.From.UTC().Equal(req.To.UTC()) {
		return errors.New("from can't be after or equal to")
	}
	return nil
}

// CreateOrderV1Response — POST v1/orders response model
type CreateOrderV1Response struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	Status    string    `json:"status"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}
