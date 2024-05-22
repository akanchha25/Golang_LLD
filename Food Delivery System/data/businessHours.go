package data

import "time"

// BusinessHours represents the business hours of a restaurant or food service.
type BusinessHours struct {
    StartTime time.Time
    EndTime   time.Time
}

// NewBusinessHours creates a new instance of BusinessHours.
func NewBusinessHours(startTime, endTime time.Time) *BusinessHours {
    return &BusinessHours{
        StartTime: startTime,
        EndTime:   endTime,
    }
}

// GetStartTime returns the start time of the business hours.
func (b *BusinessHours) GetStartTime() time.Time {
    return b.StartTime
}

// GetEndTime returns the end time of the business hours.
func (b *BusinessHours) GetEndTime() time.Time {
    return b.EndTime
}
