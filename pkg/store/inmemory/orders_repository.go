package inmemory

import (
	"math/rand"
	"time"

	"github.com/alewkinr/example-app-design-review/internal/orders"
)

type OrdersRepository struct {
	// store — in-memory storage for orders
	store map[int]*orders.Order
	rnd   *rand.Rand
}

// NewOrdersRepository — creates new instance of OrdersRepository
func NewOrdersRepository() *OrdersRepository {
	return &OrdersRepository{
		store: make(map[int]*orders.Order),
		rnd:   rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// SaveOrder — saves order to the repository
func (r *OrdersRepository) SaveOrder(o *orders.Order) (*orders.Order, error) {
	if o.ID == 0 {
		o.ID = r.rnd.Int()
	}

	r.store[o.ID] = o
	return o, nil
}
