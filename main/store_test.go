package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	if pathKey.Filename != expectedName {
		t.Errorf("Path %s want %s", pathKey.Filename, expectedOriginalKey)
	}
}
func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := newStore(opts)
	key := "special"
	data := []byte("some picture")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := newStore(opts)
	key := "special"
	data := []byte("some picture")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}
	b, _ := ioutil.ReadAll(r)
	fmt.Println(string(b))
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}
	s.Delete(key)
}
