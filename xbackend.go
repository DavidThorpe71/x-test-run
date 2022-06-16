package main

import (
	"errors"
	"fmt"
	"log"
	"tdd-practice/configFetcher"
	"tdd-practice/feedFetcher"
	"tdd-practice/itemFetcher"
	"tdd-practice/moduleFetcher"
)

type XBackend interface {
	CreateBlob() (string, error)
}

type fetcherFunc func(item string, numberOfItems int) (string, error)

type xBackend struct {
	configFetcher configFetcher.ConfigFetcher
	moduleFetcher moduleFetcher.ModuleFetcher
	itemFetcher itemFetcher.ItemFetcher
	feedFetcher feedFetcher.FeedFetcher
	fetcherMap map[string]fetcherFunc
}

func (x *xBackend) CreateBlob() (string, error) {
	config, err := x.configFetcher.GetConfig()

	if err != nil {
		return ``, errors.New(`unable to fetch the config: `+err.Error())
	}

	moduleDataMap := make(map[string][]string)

	for module, enabled := range config {
		if enabled == 1 {
			moduleData, moduleDataErr := x.moduleFetcher.GetModule(module)

			if moduleDataErr != nil {
				return "", moduleDataErr
			}
			moduleDataMap[module] = moduleData
		}
	}

	counterMap := make(map[string]int)

	for _, itemList := range moduleDataMap {
		for _, item := range itemList {
			if counterMap[item] == 0 {
				counterMap[item] = 1
			} else {
				counterMap[item] += 1
			}
		}
	}

//	for counterItem, count := range counterMap {
//		fetchFunc := x.fetcherMap[counterItem]
//		res, fetchErr := fetchFunc(counterItem, count)
//	}
	fmt.Println(moduleDataMap)
	fmt.Println(counterMap)



	return "", errors.New("no blob to return")
}

func NewXBackend(configFetcher configFetcher.ConfigFetcher, moduleFetcher moduleFetcher.ModuleFetcher) XBackend {
	if configFetcher == nil {
		log.Panic(`missing configFetcher dependency`)
	}

	if moduleFetcher == nil {
		log.Panic(`missing configFetcher dependency`)
	}

	return &xBackend{
		configFetcher: configFetcher,
		moduleFetcher: moduleFetcher,
	}
}
