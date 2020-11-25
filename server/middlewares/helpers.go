package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/onlinejudge-server/config"
)

// VerifyJWT verifies the jwt
func VerifyJWT(conf *config.Config) fiber.Handler {
	return func(c *fiber.Ctx)error {
		accessToken := c.Get("access-token")
		_, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(conf.SecretKey), nil
		})
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}

// UniqueHandles THIS IS NOT A MIDDLEWARE only a helper
// UniqueHandles only adds unique handles to the array
func UniqueHandles(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}