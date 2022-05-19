package middlewares

import (
	"github.com/0149Sailesh/iot-server/controllers/Auth"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		var l = log.WithFields(log.Fields{
			"method":"AuthorizeJWT",
		})
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := Auth.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			l.Info(claims)
		} else {
			l.Fatal(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}