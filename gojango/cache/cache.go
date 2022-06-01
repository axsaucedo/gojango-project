package cache

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int)
	Forget(string) error
	EmptyByMatch(string) error
	Empty() error
}

type RedisCache struct {
}
