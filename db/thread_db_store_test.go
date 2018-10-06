package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/sandfort/goard/core"

	_ "github.com/go-sql-driver/mysql"
)

func TestThreadStoreContract(t *testing.T) {
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	store := NewThreadDbStore(db)
	contract := core.NewThreadStoreContract(store)
	contract.Verify(t)
}
