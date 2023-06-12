package models

import "gorm.io/gorm"

// This contain the structure of the slot table
type SlotTable struct {
	gorm.Model
	Date         string `json:"date"`
	SlotNumber   int    `json:"slotNumber"`
	TimeSlot     string `json:"timeSlot"`
	NumberOfCars int    `json:"numberOfCars"`
}
