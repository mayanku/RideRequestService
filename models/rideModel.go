package models

import (
	"gorm.io/gorm"
	booking_status "riderequest/utils/enums"
	"time"
)

// Booking Structure
type RideRequest struct {
	gorm.Model

	BookingId            string                       `json:"booking_id"`
	UserId               string                       `json:"user_id"`
	BookingDate          string                       `json:"booking_date"`
	RideDate             string                       `json:"ride_date"`
	SourceAddress        string                       `json:"source_address"`
	SourceCity           string                       `json:"source_city"`
	SourceState          string                       `json:"source_state""`
	SourcePincode        int                          `json:"source_pincode"`
	SourceLongitude      float32                      `json:"source_longitude"`
	SourceLatitude       float32                      `json:"source_latitude"`
	DestinationLatitude  float32                      `json:"destination_latitude"`
	DestinationLongitude float32                      `json:"destination_longitude"`
	DestinationAddress   string                       `json:"destination_address"`
	DestinationCity      string                       `json:"destination_city"`
	DestinationState     string                       `json:"destination_state""`
	DestinationPincode   int                          `json:"destination_pincode"`
	BookingTime          string                       `json:"booking_time"`
	RideTime             time.Time                    `json:"ride_time"`
	SlotTime             string                       `json:"slot_time"`
	SlotNumber           int                          `json:"slot_number"`
	Distance             string                       `json:"distance"`
	Amount               int                          `json:"amount"`
	Status               booking_status.BookingStatus `json:"status"`
	CreatedAt            time.Time                    `json:"created_at"`
	UpdatedAt            time.Time                    `json:"updated_at"`
}
