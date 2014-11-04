package handler
import (
	compojure "compojure/core"
	"compojure/route"
        "ring/middleware/defaults"
	"ring/util/response"
	hickory "hickory/core"
	s "hickory/select"
)
import type (
	java.net.{URI, URLDecoder}
	java.io.FileNotFoundException
)

func tee(x, label) {
	println(label, "=", x)
	flush()
	x
}

func score(imgElement) {
	attrs := imgElement(ATTRS)
	width := attrs(WIDTH)
	height := attrs(HEIGHT)
	[
		if width && height {
			Float::parseFloat(width) * Float::parseFloat(height)
		} else {
			0
		},
		attrs(SRC)
	]
}

func image(url) {
	// client.get(url)(BODY)
	try {
		tree := hickory.asHickory(hickory.parse(slurp(url)))
		elements := (s.tag(IMG)  s.\select\  tree)
		if isEmpty(elements) {
			"http://upload.wikimedia.org/wikipedia/commons/a/ac/No_image_available.svg"
		} else {
			relativeUri string := last(first  sortBy  (score  map  elements))[1]
			new URI(url)->resolve(relativeUri)->toString()
		}
	} catch FileNotFoundException e {
		println(e->getMessage())
		"http://upload.wikimedia.org/wikipedia/commons/a/aa/Empty_set.svg"
	}
}


compojure.defroutes(
	appRoutes,
	compojure.GET(
		"/",
		{params: QUERY_PARAMS},
		response.redirect(image(URLDecoder::decode(params("page"), "UTF-8")))
	),
	route.notFound("Not Found")
)

var App = defaults.wrapDefaults(appRoutes, defaults.siteDefaults)
