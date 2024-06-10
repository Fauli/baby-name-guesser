package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// ValidateSession checks if a valid session is present.
// If not, it will return a 401 unauthorized error.
func ValidateSession(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {

		if !strings.EqualFold(c.Request.URL.Path, "/api/voters/login") &&
			!strings.EqualFold(c.Request.URL.Path, "/api/voters") &&
			!strings.HasPrefix(c.Request.URL.Path, "/game") {
			session, _ := store.Get(c.Request, "session")
			if session.Values["authenticated"] != true {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				fmt.Println("Unauthorized access")
				c.Abort()
			} else {
				user := session.Values["email"]
				name := session.Values["name"]
				lastName := session.Values["last_name"]
				fmt.Printf("Session is valid by user: %s %s [%s]\n", name, lastName, user)
				c.Next()
			}
		}

		c.Next()
	}

}
