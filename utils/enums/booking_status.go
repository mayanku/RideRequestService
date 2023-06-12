package booking_status

import "errors"

type BookingStatus int

const (
	PaymentPending BookingStatus = iota
	PaymentFailed
	PaymentCompleted
	RideCompleted
)

func (s BookingStatus) String() (string, error) {
	switch s {
	case PaymentPending:
		return "PaymentPending", nil
	case PaymentFailed:
		return "PaymentFailed", nil
	case PaymentCompleted:
		return "PaymentCompleted", nil
	case RideCompleted:
		return "RideCompleted", nil
	}
	return "Unknown Status", errors.New("unknown status")
}
