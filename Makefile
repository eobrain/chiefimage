# For deployment to Google App Engine
#
# USAGE:
#     make  VERSION=x.x.x  appengine-local
#     make  VERSION=x.x.x


PROJECT=chiefimage

WAR=target/$(PROJECT)-$(VERSION)-standalone.war

AE_VERSION=`echo $(VERSION) | tr .A-Z -a-z`

appengine-local: $(WAR)
	: Run on App Engine local development server
	unzip -q $(WAR) -d /tmp/exploded$$$$; dev_appserver.sh /tmp/exploded$$$$

appengine: $(WAR)
	: Deploy to App Engine
	unzip -q $(WAR) -d /tmp/exploded$$$$; appcfg.sh --oauth2 --version=$(AE_VERSION) update /tmp/exploded$$$$

$(WAR):
	lein do  fgoc --force,  ring uberwar

clean:
	rm -f $(WAR)
	lein clean
