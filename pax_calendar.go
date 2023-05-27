package calendar

import "time"

// PaxCalendar is a calendar as defined in https://en.wikipedia.org/wiki/Pax_Calendar
type PaxCalendar struct {
	t time.Time
}

type PaxMonth int

const (
	standardYearDays = 364
	leapYearDays     = 371
)

var startDate = time.Date(1928, time.January, 1, 0, 0, 0, 0, time.UTC)

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

func isLeapYear(year int) bool {
	return year%100 == 99 || year%6 == 0 || year%400 == 0
}

func daysSince(t time.Time) int {
	return int(t.Sub(startDate).Hours() / 24)
}

func fullYearsSince(t time.Time) (standardYears int, leapYears int) {
	days := daysSince(t)
	standardYears = 0
	leapYears = 0
	startYear := startDate.Year()

	for {
		if isLeapYear(startYear + leapYears + standardYears) {
			days -= leapYearDays
			if days < 0 {
				break
			}
			leapYears++
		} else {
			days -= standardYearDays
			if days < 0 {
				break
			}
			standardYears++
		}
	}
	return
}

func startOfYear(t time.Time) time.Time {
	standardYears, leapYears := fullYearsSince(t)
	daysSinceStartDate := standardYears*standardYearDays + leapYears*leapYearDays
	return startDate.Add(time.Hour * time.Duration(24*daysSinceStartDate))
}
