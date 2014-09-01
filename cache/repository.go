package cache

import "github.com/thbourlove/design-patterns/cache/store"

type Repository struct {
	storage store.Interface
}

func (c *Repository) Get(key string) (value string) {
	return c.storage.Get(key)
}

func (c *Repository) Set(key, value string) {
	c.storage.Set(key, value)
}

func NewRepository(s store.Interface) (r *Repository) {
	return &Repository{s}
}
