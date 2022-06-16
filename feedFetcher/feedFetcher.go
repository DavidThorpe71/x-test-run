package feedFetcher

import (
	"tdd-practice/itemFetcher"
)
/*
make mocks command:
mockgen -source .\configFetcher.go -destination .\configFetcher_mock.go -package configFetcher
 */

type Feed struct {
	Items []itemFetcher.Item
}

type FeedFetcher interface {
	GetFeed(feedName string) (Feed, error)
}
