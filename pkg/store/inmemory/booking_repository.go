package inmemory

import (
	"sync"
	"time"

	"github.com/alewkinr/example-app-design-review/internal/booking"
)

// BookingRepository — in-memory storage for bookings
type BookingRepository struct {
	sync.RWMutex
	// store — in-memory storage for orders
	store map[string][]booking.Booking
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		store: make(map[string][]booking.Booking),
	}
}

func (r *BookingRepository) SelectIntersectedBookings(roomID string, from, to time.Time) ([]booking.Booking, error) {
	r.RLock()
	roomBookings, ok := r.store[roomID]
	r.RUnlock()
	if !ok {
		return nil, nil
	}

	withIntersections := make([]booking.Booking, 0)
	r.RLock()
	for _, b := range roomBookings {
		if r.isIntersected(b, from, to) {
			withIntersections = append(withIntersections, b)
		}
	}
	r.RUnlock()

	return withIntersections, nil
}

// SaveBooking — saves booking to the storage
func (r *BookingRepository) SaveBooking(b booking.Booking) (booking.Booking, error) {
	r.Lock()
	r.store[b.Room.ID] = append(r.store[b.Room.ID], b)
	r.Unlock()
	return booking.Booking{}, nil
}

// isIntersected — checks if booking is intersected with other bookings within the given time range
func (r *BookingRepository) isIntersected(b booking.Booking, from, to time.Time) bool {
	f, t := from.UTC(), to.UTC()
	chin, chout := b.CheckInDateTime.UTC(), b.CheckOutDateTime.UTC()

	// |in|  [f]    |out|
	if (f.After(chin) || f.Equal(chin)) &&
		(f.Before(chout) || f.Equal(chout)) {
		return true
	}
	// |in|  [t]    |out|
	if (t.After(chin) || t.Equal(chin)) &&
		(t.Before(chout) || t.Equal(chout)) {
		return true
	}

	return false
}
