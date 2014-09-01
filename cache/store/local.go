package store

type Local struct {
	storage map[string]string
}

func (s *Local) Get(key string) (value string) {
	return s.storage[key]
}

func (s *Local) Set(key, value string) {
	s.storage[key] = value
}

func NewLocal() *Local {
	return &Local{make(map[string]string)}
}
