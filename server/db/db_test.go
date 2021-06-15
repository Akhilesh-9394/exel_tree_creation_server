package db

import "testing"

var (
	database = Newdatabase()
)

func TestDatabaseClient(t *testing.T) {
	client, ok := database.DBinstance()
	if !ok && client == nil {
		t.Errorf("Mongo failed to connect")
	}

}
