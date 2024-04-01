package booking

import "time"

type Filter struct {
	RoomID           string
	CheckInDateTime  time.Time
	CheckOutDateTime time.Time
}

// Repository — data layer for managing
type Repository interface {
	SelectIntersectedBookings(roomID string, from, to time.Time) ([]*Booking, error)
	SaveBooking(b *Booking) (*Booking, error)
}
