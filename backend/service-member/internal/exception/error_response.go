package exception

import "time"

type ErrorResponse struct {
	Message string
	Date    time.Time
}
