package integration_tests

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

const (
	DatabaseRootPassword = "Testing1"
	DatabaseName         = "test"
	DatabaseUserName     = "john"
	DatabaseUserPassword = "Testing1"
)

// EmbeddedMysql spins up a mysql container.
func EmbeddedMysql(t *testing.T, migrationsFilePath string) (host string, port int) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mysql",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": DatabaseRootPassword,
			"MYSQL_DATABASE":      DatabaseName,
			"MYSQL_USER":          DatabaseUserName,
			"MYSQL_PASSWORD":      DatabaseUserPassword,
		},
		WaitingFor: wait.ForListeningPort("3306/tcp"),
	}
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	t.Cleanup(func() {
		t.Log("terminating mysql container")
		err := mysqlC.Terminate(ctx)
		if err != nil {
			t.Fatal(err)
		}
	})

	host, err = mysqlC.Host(ctx)
	assert.NoError(t, err)
	p, err := mysqlC.MappedPort(ctx, "3306/tcp")
	assert.NoError(t, err)
	port = p.Int()

	m, err := getMigrate(host, port, migrationsFilePath)
	if err != nil {
		t.Fatal(err, "getting migration object failed")
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			//klog.Info("No changes")
		}
		t.Fatal(err, "failed migration object failed")
	}
	//klog.Info("migration successful")

	return host, port
}

func getMigrate(host string, port int, migrationsFilePath string) (*migrate.Migrate, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		DatabaseUserName, DatabaseRootPassword, host, port, DatabaseName)

	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate database config. invalid DB url")
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return migrate.NewWithDatabaseInstance(migrationsFilePath, DatabaseName, driver)
}
