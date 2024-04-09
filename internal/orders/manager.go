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
func (m *Manager) CreateOrder(o Order) (Order, error) {
	o.Status = StatusCreated

	// we save order, but we don't check if room is available
	// keep in mind that room booking checking can be done asynchronously
	newOrder, createOrderErr := m.ordersRepository.SaveOrder(o)
	if createOrderErr != nil {
		m.log.Error("save order", "order", o, "error", createOrderErr)
		return Order{}, ErrSaveOrder
	}

	b := booking.Booking{
		Room: booking.Room{
			ID:      newOrder.RoomID,
			HotelID: newOrder.HotelID,
		},
		CheckInDateTime:  newOrder.CheckInDateTime,
		CheckOutDateTime: newOrder.CheckOutDateTime,
	}

	// to simplify the example we check room availability synchronously, save the booking, and change order status
	isRoomAvailable := m.roomsManager.IsRoomAvailable(b)
	if !isRoomAvailable {
		m.log.Debug("room is not available", "booking", b)
		newOrder.Status = StatusDeclined

		if _, updateOrderErr := m.ordersRepository.SaveOrder(newOrder); updateOrderErr != nil {
			m.log.Error("update order", "order", newOrder, "booking", b, "error", createOrderErr)
			return Order{}, ErrSaveOrder
		}

		return Order{}, nil
	}

	if saveBookingErr := m.roomsManager.UpdateBooking(b); saveBookingErr != nil {
		m.log.Error("save booking", "order", newOrder, "booking", b, "error", saveBookingErr)
		return Order{}, ErrSaveOrder
	}

	newOrder.Status = StatusApproved
	approvedOrder, updateOrderErr := m.ordersRepository.SaveOrder(newOrder)
	if updateOrderErr != nil {
		m.log.Error("update order", "order", newOrder, "booking", b, "error", createOrderErr)
		return Order{}, ErrSaveOrder
	}

	return approvedOrder, nil
}
