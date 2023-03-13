package middlewares

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"DDD_Project/infrastructure/config"
	"DDD_Project/infrastructure/util"
)

func JWTAuth(config *config.AppConfig) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": "token cannot be null",
				})
				return errors.New("token cannot be null")
			}
			j := util.NewJWT(config)
			extractToken, err := j.ExtractToken(token)
			if err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})
				return err
			}
			claims, err := j.ParseToken(extractToken)
			if err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})
				return err
			}

			c.Set("claims", claims)
			return next(c)
		}
	}
}
