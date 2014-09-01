package store_test

import (
	"testing"

	"github.com/thbourlove/design-patterns/cache/store"
)

func TestLocalGetSet(t *testing.T) {
	store := store.NewLocal()
	key := "foo"
	store.Set(key, "bar")
	value := store.Get(key)
	if value != "bar" {
		t.Errorf("Value = %s; want bar", value)
	}
}
