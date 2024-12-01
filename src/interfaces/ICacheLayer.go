package interfaces

type ICacheLayer interface {
	Get(key string) (*string, error)
	Set(key string, value string) error
	Remove(key []string) error
}
