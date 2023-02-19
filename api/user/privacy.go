package user

import (
	"github.com/hiejulia/api-online-book-store/utils"
)

// Setup ...
func SetupPrivacy() {
	TokenSecret = []byte(utils.GetEnvStr("CACHE_SECRET"))
	if len(TokenSecret) == 0 {
		panic("missing CACHE_SECRET environment variable")
	}
}
