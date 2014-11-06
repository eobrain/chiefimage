package handler
import (
	"clojure/string"
	compojure "compojure/core"
	"compojure/route"
        "ring/middleware/defaults"
	"ring/util/response"
	hickory "hickory/core"
	s "hickory/select"
)
import type (
	java.net.{URI, URLDecoder, ConnectException}
	java.io.{IOException, FileNotFoundException}
)

func tee(x, label) {
	println(label, "=", x)
	flush()
	x
}

func suffix(s string) {
	dot := s->lastIndexOf(int('.'))
	string.lowerCase(s->substring(dot + 1))
}

func typeScore(imgElement, src) {
	switch suffix(src) {
	case "gif": 0.01
	default: 1
	}
}

func areaScore(imgElement, src) {
	try {
		attrs := imgElement(ATTRS)
		width := attrs(WIDTH)
		height := attrs(HEIGHT)
		if width && height {
			Float::parseFloat(width) * Float::parseFloat(height)
		} else {
			0
		}
	} catch NumberFormatException _ {
		0
	}
}

func score(imgElement) {
	src := imgElement(ATTRS)(SRC)
	[
		areaScore(imgElement, src) * typeScore(imgElement, src),
		src
	]
}

func image(url) {
	try {
		fallback := str(
			"http://api.webthumbnail.org/?width=300&height=300&url=",
			string.replace(url, /https?:\/\//, "")
		)
		tree := hickory.asHickory(hickory.parse(slurp(url)))
		elements := (s.tag(IMG)  s.\select\  tree)
		if isEmpty(elements) {
			fallback
		} else {
			[score, relativeUri string] := last(first  sortBy  (score  map  elements))
			if score < 100 {
				fallback
			} else {
				new URI(url)->resolve(relativeUri)->toString()
			}
		}
	} catch ConnectException  e {
		println(url, e->getMessage())
		"http://pixabay.com/static/uploads/photo/2012/04/24/13/00/mark-39951_640.png"
	} catch FileNotFoundException e {
		println(url, e->getMessage())
		"http://upload.wikimedia.org/wikipedia/commons/a/aa/Empty_set.svg"
	} catch IOException e {
		println(url, e->getMessage())
		"http://pixabay.com/static/uploads/photo/2012/04/24/13/00/mark-39951_640.png"
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
