package inmemory

import (
	"time"

	"github.com/alewkinr/example-app-design-review/internal/booking"
)

// BookingRepository — in-memory storage for bookings
type BookingRepository struct {
	// store — in-memory storage for orders
	store map[string][]*booking.Booking
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{store: make(map[string][]*booking.Booking)}
}

func (r *BookingRepository) SelectIntersectedBookings(roomID string, from, to time.Time) ([]*booking.Booking, error) {
	roomBookings, ok := r.store[roomID]
	if !ok {
		return nil, nil
	}

	withIntersections := make([]*booking.Booking, 0)
	for _, b := range roomBookings {
		if from.Before(b.CheckInDateTime) ||
			from.After(b.CheckOutDateTime) ||
			from.Before(b.CheckOutDateTime) ||
			to.After(b.CheckInDateTime) ||
			to.After(b.CheckOutDateTime) ||
			to.Before(b.CheckOutDateTime) {
			withIntersections = append(withIntersections, b)
		}
	}

	return withIntersections, nil
}

// SaveBooking — saves booking to the storage
func (r *BookingRepository) SaveBooking(b *booking.Booking) (*booking.Booking, error) {
	r.store[b.Room.ID] = append(r.store[b.Room.ID], b)
	return nil, nil
}
