package exception

import "time"

type BaseError struct {
	status  int
	message string
	date    time.Time
}
