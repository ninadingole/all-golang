package employees_test

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	mig "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"github.com/ninadingole/all-golang/pkg/employees"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type TestLogConsumer struct {
}

func (g *TestLogConsumer) Accept(l testcontainers.Log) {
	print(string(l.Content))
}

func TestEmployeesIntegration(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mysql:5.6",
		ExposedPorts: []string{"3306/tcp"},
		SkipReaper:   true,
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "Testing123",
			"MYSQL_DATABASE":      "test",
			"MYSQL_USER":          "john",
			"MYSQL_PASSWORD":      "Testing123",
		},
		WaitingFor: wait.ForLog("mysqld: ready for connections"),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          false,
	})

	g := TestLogConsumer{}
	if err != nil {
		t.Error(err)
	}

	err = mysqlC.Start(ctx)
	if err != nil {
		t.Error(err)
	}

	err = mysqlC.StartLogProducer(ctx)
	if err != nil {
		t.Error(err)
	}

	mysqlC.FollowOutput(&g)

	defer mysqlC.StopLogProducer()
	defer mysqlC.Terminate(ctx)

	ip, err := mysqlC.Host(ctx)
	if err != nil {
		t.Error(err)
	}

	port, err := mysqlC.MappedPort(ctx, "3306")
	if err != nil {
		t.Error(err)
	}

	dsn := fmt.Sprintf("john:Testing123@tcp(%s:%d)/test?charset=utf8", ip, port.Int())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		t.Error(err)
	}

	instance, err := mig.WithInstance(sqlDB, &mig.Config{})
	if err != nil {
		t.Error(err)
	}
	fsrc, err := (&file.File{}).Open("file://migrations")
	if err != nil {
		t.Error(err)
	}

	dbInstance, err := migrate.NewWithInstance("file", fsrc, "test", instance)
	if err != nil {
		t.Error(err)
	}

	err = dbInstance.Up()
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
