package integration_tests

import (
	"fmt"
	_ "github.com/mattes/migrate/source/file"
	"github.com/ninadingole/all-golang/pkg/employees"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestEmployeesIntegration(t *testing.T) {
	host, port := EmbeddedMysql(t, "file://../migrations")

	dsn := fmt.Sprintf("john:Testing1@tcp(%s:%d)/test?charset=utf8", host, port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	repository := employees.NewRepository(db)

	t.Run("Should store employee", func(t *testing.T) {
		e := repository.Save(employees.Employee{
			Name:          "John Doe",
			SalaryInCents: 1000000,
			Address:       "Test Address",
			DateOfJoining: time.Now(),
			DOB:           time.Date(1995, 9, 9, 0, 0, 0, 0, time.UTC)})

		assert.NoError(t, e, "shouldn't error")
	})

}
