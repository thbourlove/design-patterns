package cache_test

import (
	"testing"

	"github.com/thbourlove/cache"
)

type MockStore struct {
}

func (s *MockStore) Get(key string) (value string) {
	return "bar"
}

func (s *MockStore) Set(key, value string) {
}

func NewMockStore() *MockStore {
	return &MockStore{}
}

func TestRepositoryGetSet(t *testing.T) {
	store := NewMockStore()
	cache := cache.NewRepository(store)
	key := "foo"
	cache.Set(key, "bar")
	value := cache.Get(key)
	if value != "bar" {
		t.Errorf("Value = %s; want bar", value)
	}
}
