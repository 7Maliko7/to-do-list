package db

import (
	"context"
	"testing"
)

var storage DbStorage

func TestNew(t *testing.T) {
	ctx := context.TODO()
	url := "postgresql://postgres:postgres@localhost:5432/todolist"
	var err error
	storage, err = New(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestDbStorage_Close(t *testing.T) {
	err := storage.Close(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
