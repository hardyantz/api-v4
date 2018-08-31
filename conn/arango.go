package conn

import (
	"context"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// ArangoHandler arangodb connection handler
func ArangoHandler() (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + os.Getenv("ARANGODB_READ_HOST") + ":" + os.Getenv("ARANGODB_READ_PORT")},
	})

	if err != nil {
		return nil, err
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(os.Getenv("ARANGODB_READ_USERNAME"), os.Getenv("ARANGODB_READ_PASSWORD")),
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	db, err := c.Database(ctx, os.Getenv("ARANGODB_READ_DATABASE"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
