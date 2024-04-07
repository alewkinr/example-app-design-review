package booking

import (
	"log/slog"
)

// Manager — bookings manager
type Manager struct {
	log                *slog.Logger
	bookingsRepository Repository
}

func NewManager(log *slog.Logger, bookingsRepository Repository) *Manager {
	return &Manager{log: log, bookingsRepository: bookingsRepository}
}

// IsRoomAvailable — manager checks if booking  is available using its private logic
func (m *Manager) IsRoomAvailable(b Booking) bool {
	bookings, err := m.bookingsRepository.SelectIntersectedBookings(b.Room.ID, b.CheckInDateTime, b.CheckOutDateTime)
	if err != nil {
		m.log.Error("select bookings", "booking", b, "error", err)
		return false
	}

	if len(bookings) == 0 {
		return true
	}

	return false
}

// UpdateBooking — updates booking
func (m *Manager) UpdateBooking(b Booking) error {
	_, updBookingErr := m.bookingsRepository.SaveBooking(b)
	return updBookingErr
}
