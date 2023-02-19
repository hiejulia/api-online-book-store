package clients

var (
	_cache *Redis
	_db    *SQL
)

// Cache will return the cache connection instance.
func Cache() *Redis {
	if _cache == nil {
		panic("Trying to access API Cache without setup.")
	}
	return _cache
}

// DB returns a pointer to the database instance.
func DB() *SQL {
	if _db == nil {
		panic("Trying access the API DB without setup.")
	}
	return _db
}

// SetCache tells the API which.
func SetCache(cache *Redis) {
	_cache = cache
}

// SetDB tell the API to use a database instance.
func SetDB(db *SQL) {
	_db = db
}
