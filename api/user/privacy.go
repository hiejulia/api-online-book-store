package user

import (
	"github.com/hiejulia/api-online-book-store/api/auth"
	"github.com/hiejulia/api-online-book-store/utils"
)

// Setup ...
func SetupPrivacy() {
	auth.TokenSecret = []byte(utils.GetEnvStr("CACHE_SECRET"))
	if len(auth.TokenSecret) == 0 {
		panic("missing CACHE_SECRET environment variable")
	}
}
