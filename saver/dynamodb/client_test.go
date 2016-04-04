package dynamodb

import (
	"testing"
)

func TestClient(t *testing.T){
	client := NewClient()
	if client == nil{
		t.Error("new client failed")
	}
}
