package handler_test
import (
	"clojure/test"
        mock "ring/mock/request"
	"chiefimage.core/handler"
)

test.deftest(testApp,
	test.testing("main route", {
		response := handler.App(
			mock.request(GET, "/?page=https://www.flickr.com/photos/eob/55434370/"))
		test.is(response(STATUS) == 302)
		test.is(response(HEADERS)("Location")
			== "https://c1.staticflickr.com/1/32/55434370_ba1783d747_z.jpg")
	}),
	test.testing("not-found route", {
		response := handler.App(mock.request(GET, "/invalid"))
		test.is(response(STATUS) == 404)
	})
)
