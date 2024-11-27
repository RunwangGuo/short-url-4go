package interfaces

type IGenerateService interface {
	Generate(urls []string, expiredTs int64) (map[string]string, error)
}
