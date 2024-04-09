package booking

import "time"

// Repository — data layer for managing
type Repository interface {
	SelectIntersectedBookings(roomID string, from, to time.Time) ([]Booking, error)
	SaveBooking(b Booking) (Booking, error)
}
