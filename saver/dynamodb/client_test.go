package dynamodb

import (
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Error("new client failed")
	}
	table := "table"
	client.SetTableName(table)
	if client.table != table {
		t.Error("set table name failed")
	}
}
