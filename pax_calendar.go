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
var monthLengthsStandard = []int{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 0, 28}
var monthLengthsLeap = []int{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 7, 28}

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
	days := day - 1
	for y := startDate.Year(); y < year; y++ {
		if isLeapYear(y) {
			days += leapYearDays
		} else {
			days += standardYearDays
		}
	}
	monthLengths := monthLengthsStandard
	if isLeapYear(year) {
		monthLengths = monthLengthsLeap
	}
	for m := 0; m < int(month)-1; m++ {
		days += monthLengths[m]
	}
	t := startDate.Add(time.Hour * time.Duration(24*days))
	return &PaxCalendar{t}
}

func NewFromTime(t time.Time) *PaxCalendar {
	return &PaxCalendar{t}
}

func Now() *PaxCalendar {
	return &PaxCalendar{time.Now()}
}

func (p *PaxCalendar) Year() (year int) {
	y, _, _ := p.Date()
	return y
}

func (p *PaxCalendar) Day() (day int) {
	_, _, d := p.Date()
	return d
}

func (p *PaxCalendar) Month() (month PaxMonth) {
	_, m, _ := p.Date()
	return m
}

func (p *PaxCalendar) Date() (year int, month PaxMonth, day int) {
	s, l := fullYearsSince(p.t)
	year = startDate.Year() + s + l

	yearDay := p.YearDay()
	month = 1
	day = 1
	monthLengths := monthLengthsStandard
	if isLeapYear(year) {
		monthLengths = monthLengthsLeap
	}
	for {
		yearDay -= monthLengths[month-1]
		if yearDay < 0 {
			day = yearDay + monthLengths[month-1] + 1
			break
		}
		month++
	}

	return year, month, day
}

func (p *PaxCalendar) YearDay() (day int) {
	return int(p.t.Sub(startOfYear(p.t)).Hours() / 24)
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return false
	}
	lastTwo := year % 100

	return lastTwo%6 == 0 || lastTwo == 0 || lastTwo == 99
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
	return standardYears, leapYears
}

func startOfYear(t time.Time) time.Time {
	standardYears, leapYears := fullYearsSince(t)
	daysSinceStartDate := standardYears*standardYearDays + leapYears*leapYearDays
	return startDate.Add(time.Hour * time.Duration(24*daysSinceStartDate))
}
