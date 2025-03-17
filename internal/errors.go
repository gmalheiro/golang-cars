package internal

import "errors"

var ErrNotAvailableCars = errors.New("there are no available cars")
var ErrFieldsNotPropperlyField = errors.New("invalid field")
var ErrExistingItem = errors.New("existing item")
var ErrItemNotFound = errors.New("not found item")
