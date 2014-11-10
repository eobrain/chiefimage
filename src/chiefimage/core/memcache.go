package memcache
import type com.google.appengine.api.memcache.MemcacheServiceFactory

cache := MemcacheServiceFactory::getMemcacheService()


// Return a memoized version of a function of one argument
func Memoized(f) {
	func(key) {
		cached := cache->get(key)
		if cached == nil {
			newValue := f(key)
			cache->put(key, newValue)
			newValue
		} else {
			cached
		}
	}
}
