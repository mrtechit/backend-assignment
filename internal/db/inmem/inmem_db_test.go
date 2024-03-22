package inmem

import (
	"bytes"
	"errors"
	"testing"
)

func TestDatabaseGetSuccess(t *testing.T) {
	db := New()
	value := []byte("value")
	err := db.Put("test", [][]byte{value})
	if err != nil {
		t.Fatalf("error in put operation")
	}
	b, err := db.Get("test")
	if err != nil {
		t.Fatalf("error in get operation")
	}
	if !bytes.Equal(b[0], value) {
		t.Fatalf("error in get operation")
	}
}

func TestDatabaseGetFailed(t *testing.T) {
	db := New()
	notFounndErr := errors.New("not found")
	_, err := db.Get("test")
	if err.Error() != notFounndErr.Error() {
		t.Fatalf("error in get operation")
	}
}
