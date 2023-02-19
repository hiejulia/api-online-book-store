package user

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/hiejulia/api-online-book-store/api/common"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

// Constants
var (
	TokenDuration = 24 * time.Hour
	TokenIssuer   = "api.brickbanker.fans"
	TokenSecret   []byte

	TimeoutIgnore = 5 * time.Minute
	TimeoutLong   = 10 * time.Second
	TimeoutShort  = 3 * time.Second
)

var CIDRBanList []*net.IPNet
var IPBanList []string

var ignoreRoutes = []string{
	"/privacy/data",
}

var slowRoutes = []string{
	"/pub/login",
	"/pub/register",
}

var ErrMissingAuthorization = fmt.Errorf("missing Authorization header")
var ErrInvalidTokenAlgorithm = fmt.Errorf("invalid authorization token algorithm")
var ErrInvalidTokenClaims = fmt.Errorf("invalid authorization token claims")
var ErrInvalidToken = fmt.Errorf("invalid authorization token")
var ErrForbidden = fmt.Errorf("You can't access")

// Authorize ...
func Authorize(c *gin.Context) {
	header := Extract(c)
	if header == "" {
		common.Error(c, http.StatusUnauthorized, ErrMissingAuthorization)
		return
	}

	token, err := jwt.ParseWithClaims(header, &MainClaims{}, func(tok *jwt.Token) (interface{}, error) {
		if tok.Method.Alg() != jwt.SigningMethodHS512.Name {
			return nil, ErrInvalidTokenAlgorithm
		}
		return TokenSecret, nil
	})
	if err != nil {
		common.Error(c, http.StatusUnauthorized, err)
		return
	}

	claims, ok := token.Claims.(*MainClaims)
	if !ok {
		fmt.Println("claims not ok", claims, token.Claims.(MainClaims))
		common.Error(c, http.StatusUnauthorized, ErrInvalidTokenClaims)
		return
	} else if !token.Valid {
		common.Error(c, http.StatusUnauthorized, ErrInvalidToken)
		return
	}

	// Validate the claims data.
	now := time.Now().UTC().Unix()
	if claims.Audience != "mobile" {
		err = fmt.Errorf("token has incorrect audience")
	} else if claims.ExpiresAt < now {
		err = fmt.Errorf("token has expired")
	} else if claims.IssuedAt > now {
		err = fmt.Errorf("token cannot be issued in the future")
	} else if claims.Issuer != TokenIssuer {
		err = fmt.Errorf("token not issued by API")
	} else if claims.NotBefore > now {
		err = fmt.Errorf("token active in the future")
	} else if claims.Subject != "" {
		// Do nothing here without the builder name.
	} else if (claims.ExpiresAt - claims.IssuedAt) != (TokenDuration.Milliseconds() / 1000) {
		err = fmt.Errorf("token duration incorrect")
	}
	if err != nil {
		fmt.Println("ERROR:", err)
		common.Error(c, http.StatusUnauthorized, ErrInvalidTokenClaims)
		return
	}

	// Validate that the token is found in the cache.
	var builderID string
	cache := clients.Cache()
	if builderID, err = cache.Get(header); err != nil {
		common.Error(c, http.StatusUnauthorized, ErrInvalidToken)
		return
	}

	// Validate that the builder ids match.
	if builderID != claims.Id {
		fmt.Println(builderID, "!=", claims.Id)
		common.Error(c, http.StatusUnauthorized, ErrInvalidTokenClaims)
		return
	}

	// Load the builder from the clients.
	builder := &models.User{ID: claims.Id}
	if err := clients.DB().First(builder); err != nil {
		fmt.Println("unable to find builder", claims.Id)
		common.Error(c, http.StatusUnauthorized, ErrInvalidTokenClaims)
		return
	}

	// Attach the builder to the request and continue the chain.
	c.Set("builder", builder)
	c.Next()
}

// AuthorizeAdmin ...
//func AuthorizeAdmin(c *gin.Context) {
//	builder := c.MustGet("builder").(*models.User)
//	if !builder.HasRole(constants.TOKEN_ROLE_MOD) {
//		//Error(c, http.StatusForbidden, ErrNoPermission)
//		//return
//	}
//	c.Next()
//}

// Database will attach a clients instance to the context.
func Database(ignore, slow []string) func(*gin.Context) {
	return func(c *gin.Context) {
		dur := TimeoutShort
		found := false
		// Ignore routes have 5 minutes.
		for _, ig := range ignore {
			if c.Request.URL.Path == ig {
				found = true
				break
			}
		}
		if found {
			dur = TimeoutIgnore
		} else {
			// Slow routes have 10 seconds.
			for _, ig := range slow {
				if c.Request.URL.Path == ig {
					found = true
					break
				}
			}
			if found {
				dur = TimeoutLong
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), dur)
		defer cancel()

		db := clients.DB().WithContext(ctx)
		defer db.Close()

		c.Set("db", db)
		c.Next()
	}
}

// Extract JWT from Authorization header.
func Extract(c *gin.Context) string {
	s := c.Request.Header.Get("Authorization")
	return removeBearer(s)
}

func removeBearer(s string) string {
	s = strings.Replace(s, "bearer ", "", 1)
	s = strings.Replace(s, "Bearer ", "", 1)
	return s
}

type ReqWithEmail struct {
	Email       string `json:"email"`
	BuilderName string `json:"builder_name"`
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Code returns a one-time password sent in email and used
// as verification or 2FA.
func Code(chars int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, chars)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type MainClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

// Token ...
func Token(builder *models.User, role string) (string, error) {
	date := time.Now().UTC()

	claims := &MainClaims{
		jwt.StandardClaims{
			Audience:  "mobile",
			ExpiresAt: date.Add(TokenDuration).Unix(),
			Id:        builder.ID,
			IssuedAt:  date.Unix(),
			Issuer:    TokenIssuer,
			NotBefore: date.Unix(),
			Subject:   builder.Email,
		},
		role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(TokenSecret)
	if err != nil {
		fmt.Println("utils.Token tokenString", err)
		return "", err
	}

	cache := clients.Cache()
	if err = cache.Set(tokenString, builder.ID, TokenDuration); err != nil {
		fmt.Println("utils.Token cache.Set", err)
		return "", err
	}
	return tokenString, nil
}

// Setup ...
func SetupMiddleware(r *gin.Engine) {
	TokenSecret = []byte(utils.GetEnvStr("CACHE_SECRET"))
	if len(TokenSecret) == 0 {
		panic("missing CACHE_SECRET environment variable")
	}

	// Ban IP Middleware
	r.Use(func(c *gin.Context) {
		clientIpStr := c.ClientIP()
		cacheKey := "ban_" + clientIpStr
		cache := clients.Cache()
		cacheTime := time.Minute
		if cacheVal, err := cache.GetInt(cacheKey); err == nil && cacheVal == 1 {
			common.Error(c, http.StatusForbidden, ErrForbidden)
			return
		}
		cache.SetInt(cacheKey, 0, cacheTime)
	})

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		if strings.HasPrefix(param.Path, "/metrics") || strings.HasPrefix(param.Path, "/state") {
			return ""
		}
		authToken := removeBearer(param.Request.Header.Get("Authorization"))
		authTokenSplit := strings.Split(authToken, ".")
		if len(authTokenSplit) > 2 {
			if jwtStr, err := base64.RawStdEncoding.DecodeString(authTokenSplit[1]); err == nil {
				var claims jwt.StandardClaims
				json.Unmarshal(jwtStr, &claims)
				authToken = claims.Audience + "-" + claims.Subject
			}
		}
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" >%s< %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			authToken,
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"DELETE", "GET", "OPTIONS", "POST", "PUT"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))

	r.Use(Database(ignoreRoutes, slowRoutes))
}
