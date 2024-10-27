package httptransfer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/li-zeyuan/common-go/utils"
)

var UIdCtxKey = "uid"

func NotStrictAuthorizationMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, &utils.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.Next()
			return
		}

		claims, ok := token.Claims.(*utils.JwtClaims)
		if ok && token.Valid {
			if claims.Uid > 0 {
				c.Set(UIdCtxKey, claims.Uid)
				//mylogger.Infof(c.Request.Context(), "uid: %d", claims.Uid)
			}
		}

		c.Next()
	}
}

func StrictAuthorizationMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, &utils.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			ErrJSONResp(c, http.StatusForbidden, err)
			return
		}

		claims, ok := token.Claims.(*utils.JwtClaims)
		if ok && token.Valid {
			//mylogger.Infof(c.Request.Context(), "uid: %d", claims.Uid)
		} else {
			ErrJSONResp(c, http.StatusForbidden, err)
			return
		}

		c.Set(UIdCtxKey, claims.Uid)
		c.Header(UIdCtxKey, strconv.FormatInt(claims.Uid, 10))
		c.Next()
	}
}

func GetUid(c *gin.Context) int64 {
	val, ok := c.Get(UIdCtxKey)
	if !ok {
		//mylogger.Warn(c.Request.Context(), "can not transfer UIdCtxKey")
		return 0
	}

	return val.(int64)
}
