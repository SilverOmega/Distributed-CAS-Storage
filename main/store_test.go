package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "this picture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "d5b2eaa9093512bfd60d8dce9302ce095436a7a8"
	expectedName := "d5b2e/aa909/3512b/fd60d/8dce9/302ce/09543/6a7a8"
	if pathKey.Pathname != expectedName {
		t.Errorf("Path %s want %s", pathKey.Pathname, expectedName)
	}
	if pathKey.OriginalKey != expectedName {
		t.Errorf("Path %s want %s", pathKey.OriginalKey, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := newStore(opts)
	data := bytes.NewReader([]byte("some picture"))
	if err := s.writeStream("My file", data); err != nil {
		t.Error(err)
	}
}
