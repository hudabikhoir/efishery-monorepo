package middleware

import (
	"efishery/api/common"
	serviceAuth "efishery/business/auth"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var (
	jwtSigningMethod = jwt.SigningMethodHS256
)

func JWTMiddleware(auth serviceAuth.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if strings.Contains(c.Request().URL.Path, "/login") {
				return next(c)
			}

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			fmt.Println(signature)
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.NewMissingHeaderResponse())
			}
			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.NewInternalServerErrorResponse())
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(token *jwt.Token) (interface{}, error) {
				return []byte("<screet-key>"), nil
			})
			fmt.Println(c.Request().URL.Path)

			user, err := auth.Validate(signature[1])
			fmt.Println(user)
			if err != nil || user.Role == 0 {
				return c.JSON(http.StatusForbidden, common.NewExpiredTokenErrorResponse())
			}

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSigningMethod {
				return c.JSON(http.StatusForbidden, common.NewInternalServerErrorResponse())
			}

			return next(c)
		}
	}
}
