package orders

import "time"

const (
	// StatusCreated — every new order has this status
	StatusCreated = "NEW"
	// StatusApproved — room in the order is available and was booked during this order
	StatusApproved = "APPROVED"
	// StatusDeclined — room in the order is not available or got some other issues
	StatusDeclined = "DECLINED"
)

// Order — customer order model
type Order struct {
	ID        int
	UserEmail string

	HotelID string
	RoomID  string

	CheckInDateTime  time.Time
	CheckOutDateTime time.Time

	Status string
}
