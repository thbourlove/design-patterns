package store

type Interface interface {
	Get(key string) (value string)
	Set(key, value string)
}
