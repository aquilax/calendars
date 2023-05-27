package calendar

import "time"

// PaxCalendar is a calendar as defined in https://en.wikipedia.org/wiki/Pax_Calendar
type PaxCalendar struct {
	t time.Time
}

type PaxMonth int

const StartYear = 1928

const (
	January PaxMonth = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	Columbus
	Pax
	December
)

func New(year int, month PaxMonth, day int) *PaxCalendar {
	// TODO: implementation
	return &PaxCalendar{}
}

func NewFromTime(t time.Time) *PaxCalendar {
	return &PaxCalendar{t}
}

func Now() *PaxCalendar {
	return &PaxCalendar{time.Now()}
}

func (p *PaxCalendar) Year() (year int) {
	// TODO: implementation
	return 0
}

func (p *PaxCalendar) Month() (month PaxMonth) {
	// TODO: implementation
	return 0
}

func (p *PaxCalendar) Date() (year int, month PaxMonth, day int) {
	// TODO: implementation
	return 0, 0, 0
}

func (p *PaxCalendar) YearDay() (day int) {
	// TODO: implementation
	return 0
}

func (p *PaxCalendar) IsLeap() bool {
	// TODO: implementation
	return false
}
