(defproject chiefimage "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :min-lein-version "2.0.0"
  :dependencies [[org.clojure/clojure "1.6.0"]
                 [org.eamonn.funcgo/funcgo-lein-plugin "0.5.1"]
                 [clj-http "1.0.1"]
                 [hickory "0.5.4"]
                 [compojure "1.2.0"]
                 [ring/ring-defaults "0.1.2"]]
  :plugins [[org.eamonn.funcgo/funcgo-lein-plugin "0.5.1"]
            [lein-ring "0.8.13"]]
  :ring {:handler chiefimage.core.handler/App}
  :profiles
  {:dev {:dependencies [[javax.servlet/servlet-api "2.5"]
                        [ring-mock "0.1.5"]]}})
