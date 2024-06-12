package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"sbebe.ch/baby-name-guesser/pkg/utils"
)

// ValidateSession checks if a valid session is present.
// If not, it will return a 401 unauthorized error.
func ValidateSession(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.Logger.Sugar().Debugf("URL: [%s], METHOD: [%s], HEADER: [%s], BODY: [%s]", c.Request.URL, c.Request.Method, c.Request.Header, c.Request.Body)

		if !strings.EqualFold(c.Request.URL.Path, "/api/voters/login") &&
			!strings.EqualFold(c.Request.URL.Path, "/api/voters") &&
			!strings.HasPrefix(c.Request.URL.Path, "/game") &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/swagger") {
			session, _ := store.Get(c.Request, "session")
			if session.Values["authenticated"] != true {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				utils.Logger.Warn("Unauthorized access")
				c.Abort()
			} else {
				user := session.Values["email"]
				name := session.Values["name"]
				lastName := session.Values["last_name"]
				utils.Logger.Sugar().Debugf("Session is valid by user: %s %s [%s]", name, lastName, user)
				c.Next()
			}
		}

		c.Next()
	}

}
