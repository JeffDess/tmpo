package storage

import (
	"math"
	"time"
)

type TimeEntry struct {
	ID int64
	ProjectName string
	StartTime time.Time
	EndTime *time.Time
	Description string
	HourlyRate *float64
	MilestoneName *string
}

func (t *TimeEntry) Duration() time.Duration {
	if( t.EndTime == nil) {
		return time.Since(t.StartTime)
	}

	return t.EndTime.Sub(t.StartTime)
}

func (t *TimeEntry) IsRunning() bool {
	return t.EndTime == nil
}

// RoundedHours returns duration in hours rounded to 2 decimal places for billing.
// Could be made configurable to support different rounding increments (0.1h, 0.25h, etc).
func (t *TimeEntry) RoundedHours() float64 {
	return math.Round(t.Duration().Hours()*100) / 100
}

type Milestone struct {
	ID          int64
	ProjectName string
	Name        string
	StartTime   time.Time
	EndTime     *time.Time
}

func (m *Milestone) IsActive() bool {
	return m.EndTime == nil
}

func (m *Milestone) Duration() time.Duration {
	if m.EndTime == nil {
		return time.Since(m.StartTime)
	}
	return m.EndTime.Sub(m.StartTime)
}