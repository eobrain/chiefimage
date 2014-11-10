(defproject chiefimage "0.1.4"
  :description "A web service that redirects to the main image of a web page."
  :url "http://example.com/FIXME"
  :min-lein-version "2.0.0"
  :dependencies [[org.clojure/clojure "1.6.0"]
                 [org.eamonn.funcgo/funcgo-lein-plugin "0.5.1"]
                 [clj-http "1.0.1"]
                 [hickory "0.5.4"]
                 [ring "1.2.1"]
                 [compojure "1.1.8"]
                 [ring/ring-defaults "0.1.2"]
                 [com.google.appengine/appengine-api-1.0-sdk "1.9.15"]]
  :plugins [[org.eamonn.funcgo/funcgo-lein-plugin "0.5.1"]
            [lein-ring "0.8.13"]]
  :ring {:handler chiefimage.core.handler/App}
  :profiles
  {:dev {:dependencies [[javax.servlet/servlet-api "2.5"]
                        [ring-mock "0.1.5"]]}})
