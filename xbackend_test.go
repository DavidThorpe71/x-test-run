package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"tdd-practice/configFetcher"
	"tdd-practice/moduleFetcher"
	"testing"
)

func TestXBackend(t *testing.T) {
	expectedErr := errors.New(`oh no an error`)

	Convey("GIVEN we instantiate the X backend with a nil configFetcher", t, func() {
		Convey("THEN we should panic", func() {
			So(func() {
				NewXBackend(nil, nil)
			}, ShouldPanicWith, `missing configFetcher dependency`)

		})
	})

	Convey("GIVEN we instantiate a new backend with a config fetcher", t, func() {
		mockController := gomock.NewController(t)
		mockConfigFetcher := configFetcher.NewMockConfigFetcher(mockController)
		mockModuleFetcher := moduleFetcher.NewMockModuleFetcher(mockController)

		sut := NewXBackend(mockConfigFetcher, mockModuleFetcher)

		Convey("WHEN we call the CreateBlob() method", func() {
			configFetcherCall := mockConfigFetcher.EXPECT().GetConfig()

			Convey("BUT the configFetcher returns an error", func() {
				configFetcherCall.Return(configFetcher.Config{}, expectedErr)

				res, err := sut.CreateBlob()

				Convey("THEN the error is returned", func() {
					So(res, ShouldEqual, ``)
					So(err, ShouldBeError)
					So(err.Error(), ShouldEqual, `unable to fetch the config: oh no an error`)
				})
			})

			Convey("AND the configFetcher returns a Config item", func() {
				configItem := configFetcher.Config{
					"GenericEditorialModule": 1,
					"AnotherModule": 1,
					"AndAnotherModule": 1,
				}

				configFetcherCall.Return(configItem, nil)

				Convey("AND we attempt to fetch the first enabled config item's contents", func() {
					gemModuleFetcherCall := mockModuleFetcher.EXPECT().GetModule("GenericEditorialModule")

					Convey("BUT we fail to fetch the gem module data", func() {
						gemModuleFetcherCall.Return([]string{}, expectedErr)

						Convey("THEN an error is returned", func() {
							_, err := sut.CreateBlob()

							So(err, ShouldBeError)
							So(err.Error(), ShouldEqual, "oh no an error")
						})
					})

					Convey("AND we successfully fetch the GE module data", func() {
						geModuleRes := []string{"news-feed", "sports-feed", "custom-item", "news-feed"}

						gemModuleFetcherCall.Return(geModuleRes, nil)

						Convey("AND we successfully fetch the GE module data", func() {
							anotherModuleRes := []string{"news-feed", "sports-feed", "custom-item-3", "news-feed"}

							mockModuleFetcher.EXPECT().GetModule("AnotherModule").Return(anotherModuleRes, nil)

							Convey("AND we successfully fetch the AnotherModule module data", func() {
								andAnotherModuleRes := []string{"news-feed", "sports-feed", "custom-item-2", "news-feed"}

								mockModuleFetcher.EXPECT().GetModule("AndAnotherModule").Return(andAnotherModuleRes, nil)

								Convey("THEN an error is returned", func() {
									_, err := sut.CreateBlob()

									So(err, ShouldBeError)
									So(err.Error(), ShouldEqual, "no blob to return")
								})
							})
						})
					})
				})
			})
		})
	})
}






