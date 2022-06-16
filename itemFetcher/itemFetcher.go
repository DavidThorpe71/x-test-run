package itemFetcher

/*
make mocks command:
mockgen -source .\configFetcher.go -destination .\configFetcher_mock.go -package configFetcher
 */

type Item struct {
	Title string
	Description string
}

type ItemFetcher interface {
	GetItem(itemName string) (Item, error)
}
