package floxy

import (
	"errors"
)

var (
	ErrEntityNotFound   = errors.New("entity not found")
	ErrLockNotAvailable = errors.New("lock not available")
)
