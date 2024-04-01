package orders

import (
	"log/slog"

	"github.com/alewkinr/example-app-design-review/internal/booking"
)

type Manager struct {
	log              *slog.Logger
	roomsManager     *booking.Manager
	ordersRepository Repository
}

// NewManager — constructor for Manager
func NewManager(log *slog.Logger, roomsManager *booking.Manager, ordersRepository Repository) *Manager {
	return &Manager{
		log:              log,
		roomsManager:     roomsManager,
		ordersRepository: ordersRepository,
	}
}

// CreateOrder — creates order for future processing
func (m *Manager) CreateOrder(o *Order) (*Order, error) {
	// we save order, but we don't check if room is available
	// keep in mind that room booking checking can be done asynchronously
	orderWithID, createOrderErr := m.ordersRepository.SaveOrder(o)
	if createOrderErr != nil {
		m.log.Error("save order", "order", o, "error", createOrderErr)
		return nil, ErrSaveOrder
	}

	b := &booking.Booking{
		Room: &booking.Room{
			ID:      o.RoomID,
			HotelID: o.HotelID,
		},
		CheckInDateTime:  o.CheckInDateTime,
		CheckOutDateTime: o.CheckOutDateTime,
	}

	// to simplify the example we check room availability synchronously, save the booking, and change order status
	isRoomAvailable := m.roomsManager.IsRoomAvailable(b)
	if !isRoomAvailable {
		m.log.Debug("room is not available", "booking", b)
		o.Status = StatusDeclined

		if _, updateOrderErr := m.ordersRepository.SaveOrder(o); updateOrderErr != nil {
			m.log.Error("update order", "order", o, "booking", b, "error", createOrderErr)
			return nil, ErrSaveOrder
		}

		return nil, nil
	}

	if saveBookingErr := m.roomsManager.UpdateBooking(b); saveBookingErr != nil {
		m.log.Error("save booking", "order", o, "booking", b, "error", saveBookingErr)
		return nil, ErrSaveOrder
	}

	o.Status = StatusApproved
	if _, updateOrderErr := m.ordersRepository.SaveOrder(o); updateOrderErr != nil {
		m.log.Error("update order", "order", o, "booking", b, "error", createOrderErr)
		return nil, ErrSaveOrder
	}

	return orderWithID, nil
}
