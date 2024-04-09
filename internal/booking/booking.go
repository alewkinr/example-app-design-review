package booking

import "time"

// Room — model of room in hotel
type Room struct {
	ID      string
	HotelID string
}

// Booking — model of room booking in hotel
type Booking struct {
	Room Room

	CheckInDateTime  time.Time
	CheckOutDateTime time.Time
}
