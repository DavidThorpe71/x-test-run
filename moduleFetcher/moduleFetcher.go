package moduleFetcher

import "errors"

type ModuleFetcher interface {
	GetModule(moduleName string) ([]string, error)
}

type moduleFetcher struct {

}

func (m moduleFetcher) GetModule(moduleName string) ([]string, error) {
	return []string{}, errors.New(`no data for this module`)
}

func NewModuleFetcher() ModuleFetcher {

	return &moduleFetcher{}
}
