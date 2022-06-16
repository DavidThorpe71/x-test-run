package configFetcher

import "errors"
/*
make mocks command:
mockgen -source .\configFetcher.go -destination .\configFetcher_mock.go -package configFetcher
 */

type Config map[string]int

type ConfigFetcher interface {
	GetConfig() (Config, error)
}

type erroringConfigFetcher struct {

}

func (e erroringConfigFetcher) GetConfig() (Config, error) {
	return Config{}, errors.New("bang")
}

func NewErroringConfigFetcher() ConfigFetcher {
	return &erroringConfigFetcher{}
}
