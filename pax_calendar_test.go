package calendar

import (
	"reflect"
	"testing"
	"time"
)

func date(year int, month time.Month, date int) time.Time {
	return time.Date(year, month, date, 0, 0, 0, 0, time.UTC)
}

func TestNew(t *testing.T) {
	type args struct {
		year  int
		month PaxMonth
		day   int
	}
	tests := []struct {
		name string
		args args
		want *PaxCalendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want *PaxCalendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNow(t *testing.T) {
	tests := []struct {
		name string
		want *PaxCalendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Now(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaxCalendar_Year(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		wantYear int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaxCalendar{
				t: tt.fields.t,
			}
			if gotYear := p.Year(); gotYear != tt.wantYear {
				t.Errorf("PaxCalendar.Year() = %v, want %v", gotYear, tt.wantYear)
			}
		})
	}
}

func TestPaxCalendar_Month(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name      string
		fields    fields
		wantMonth PaxMonth
	}{
		{
			"1928-01-01",
			fields{date(1928, time.January, 1)},
			January,
		},
		{
			"1928-01-28",
			fields{date(1928, time.January, 28)},
			January,
		},
		{
			"1928-01-29",
			fields{date(1928, time.January, 29)},
			February,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaxCalendar{
				t: tt.fields.t,
			}
			if gotMonth := p.Month(); gotMonth != tt.wantMonth {
				t.Errorf("PaxCalendar.Month() = %v, want %v", gotMonth, tt.wantMonth)
			}
		})
	}
}

func TestPaxCalendar_Date(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name      string
		fields    fields
		wantYear  int
		wantMonth PaxMonth
		wantDay   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaxCalendar{
				t: tt.fields.t,
			}
			gotYear, gotMonth, gotDay := p.Date()
			if gotYear != tt.wantYear {
				t.Errorf("PaxCalendar.Date() gotYear = %v, want %v", gotYear, tt.wantYear)
			}
			if gotMonth != tt.wantMonth {
				t.Errorf("PaxCalendar.Date() gotMonth = %v, want %v", gotMonth, tt.wantMonth)
			}
			if gotDay != tt.wantDay {
				t.Errorf("PaxCalendar.Date() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
		})
	}
}

func TestPaxCalendar_YearDay(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantDay int
	}{
		// TODO: Add test cases.
		{
			"1928-01-01",
			fields{date(1928, time.January, 1)},
			0,
		},
		{
			"1929-01-01",
			fields{date(1929, time.January, 1)},
			2,
		},
		{
			"2023-05-25",
			fields{date(2023, time.May, 27)},
			153,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaxCalendar{
				t: tt.fields.t,
			}
			if gotDay := p.YearDay(); gotDay != tt.wantDay {
				t.Errorf("PaxCalendar.YearDay() = %v, want %v", gotDay, tt.wantDay)
			}
		})
	}
}

func Test_isLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.year); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_daysSince(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1928-01-01",
			args{date(1928, time.January, 1)},
			0,
		},
		{
			"1929-01-01",
			args{date(1929, time.January, 1)},
			366,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysSince(tt.args.t); got != tt.want {
				t.Errorf("daysSince() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fullYearsSince(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name              string
		args              args
		wantStandardYears int
		wantLeapYears     int
	}{
		{
			"1928-01-01",
			args{date(1928, time.January, 1)},
			0,
			0,
		},
		{
			"1929-01-01",
			args{date(1929, time.January, 1)},
			1,
			0,
		},
		{
			"2023-05-25",
			args{date(2023, time.May, 27)},
			79,
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStandardYears, gotLeapYears := fullYearsSince(tt.args.t)
			if gotStandardYears != tt.wantStandardYears {
				t.Errorf("fullYearsSince() gotStandardYears = %v, want %v", gotStandardYears, tt.wantStandardYears)
			}
			if gotLeapYears != tt.wantLeapYears {
				t.Errorf("fullYearsSince() gotLeapYears = %v, want %v", gotLeapYears, tt.wantLeapYears)
			}
		})
	}
}

func Test_startOfYear(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			"1928-01-01",
			args{date(1928, time.January, 1)},
			date(1928, time.January, 1),
		},
		{
			"1928-03-03",
			args{date(1928, time.March, 3)},
			date(1928, time.January, 1),
		},
		{
			"2023-05-25",
			args{date(2023, time.May, 27)},
			date(2022, time.December, 25),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startOfYear(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("startOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaxCalendar_Day(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantDay int
	}{
		{
			"1928-01-01",
			fields{date(1928, time.January, 1)},
			1,
		},
		{
			"1928-01-28",
			fields{date(1928, time.January, 28)},
			28,
		},
		{
			"1928-01-29",
			fields{date(1928, time.January, 29)},
			1,
		},
		{
			"1928-03-03",
			fields{date(1928, time.March, 3)},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaxCalendar{
				t: tt.fields.t,
			}
			if gotDay := p.Day(); gotDay != tt.wantDay {
				t.Errorf("PaxCalendar.Day() = %v, want %v", gotDay, tt.wantDay)
			}
		})
	}
}
