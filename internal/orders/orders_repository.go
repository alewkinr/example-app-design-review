package orders

// Repository — data layer for managing
type Repository interface {
	SaveOrder(o Order) (Order, error)
}
