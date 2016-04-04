package main

import (
	"github.com/orvice/password-cat/saver/dynamodb"
)

var (
	client = dynamodb.NewClient()
)

func Init() {
	client.SetTableName("")
}
