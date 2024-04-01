package orders

import "errors"

// ErrSaveOrder — general error for errors in creating order
var ErrSaveOrder = errors.New("order can't be saved")
