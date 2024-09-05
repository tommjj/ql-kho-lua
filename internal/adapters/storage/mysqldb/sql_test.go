package mysqldb

import "testing"

func TestConnection(t *testing.T) {
	_, err := NewMysqlDB()
	if err != nil {
		t.Fatal(err)
	}
}
