package interfaces

type ICacheLayer interface {
	Remove(keys []string) error
	Set(key string, value interface{}) error
	//Get(key string) (interface{}, error)
	Get(key string) (string, error)
}
