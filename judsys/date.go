package judsys

import (
	"fmt"
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

func (this Date) Year() int {
	return this.year
}

func (this Date) Month() int {
	return this.month
}

func (this Date) Day() int {
	return this.day
}

func (this Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", this.year, this.month, this.day)
}

func (this Date) ToTime(loc *time.Location) time.Time {
	return time.Date(this.year, time.Month(this.month), this.day, 0, 0, 0, 0, loc)
}
