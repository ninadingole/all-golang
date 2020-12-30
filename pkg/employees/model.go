package employees

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model

	Name          string
	DOB           time.Time
	Address       string
	Department    string
	DateOfJoining time.Time
	SalaryInCents int64
}
