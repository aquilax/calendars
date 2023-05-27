// pax implement functions to work with Calendar
package pax

import "time"

// Calendar is a calendar proposal as defined in https://en.wikipedia.org/wiki/Pax_Calendar
type Calendar struct {
	t time.Time
}

// Month represents a Calendar month
type Month int

const (
	standardYearDays = 364
	leapYearDays     = 371
)

var startDate = time.Date(1928, time.January, 1, 0, 0, 0, 0, time.UTC)
var monthLengthsStandard = []int{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 0, 28}
var monthLengthsLeap = []int{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 7, 28}

const (
	January Month = 1 + iota
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

// NewCalendar returns new Calendar, given Pax date
func New(year int, month Month, day int) *Calendar {
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
	return &Calendar{t}
}

// NewCalendarFromTime creates new Calendar from Gregorian time
func NewFromTime(t time.Time) *Calendar {
	return &Calendar{t}
}

// Year returns the Calendar year
func (p *Calendar) Year() (year int) {
	y, _, _ := p.Date()
	return y
}

// Day returns the Calendar day of the month
func (p *Calendar) Day() (day int) {
	_, _, d := p.Date()
	return d
}

// Month returns the Calendar month
func (p *Calendar) Month() (month Month) {
	_, m, _ := p.Date()
	return m
}

// Date returns the current Calendar year, month and day
func (p *Calendar) Date() (year int, month Month, day int) {
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

// YearDay returns the Calendar day of the year
func (p *Calendar) YearDay() (day int) {
	return int(p.t.Sub(startOfYear(p.t)).Hours() / 24)
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return false
	}
	lastTwo := year % 100

	return lastTwo%6 == 0 || lastTwo == 0 || lastTwo == 99
}

func (p *Calendar) IsLeapYear() bool {
	return isLeapYear(p.t.Year())
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
