package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/sandfort/goard/core"
)

func TestPostStoreContract(t *testing.T) {
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
	db, _ := sql.Open("mysql", dsn)

	postStore := NewPostDbStore(db)
	threadStore := NewThreadDbStore(db)
	contract := core.NewPostStoreContract(postStore, threadStore)
	contract.Verify(t)
}
