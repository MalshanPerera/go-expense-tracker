package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/utils"
	"github.com/golang-jwt/jwt/v5"
)

type key string

const (
	UserIDKey key = "user_id"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This should support web and mobile clients
		// Check if the user is authenticated

		// Check the Authorization header
		headerStr := r.Header.Get("Authorization")

		// If the token is not in the Authorization header, check the cookies
		if headerStr == "" {
			cookie, err := r.Cookie("token")
			if err != nil {
				utils.WriteError(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}
			headerStr = cookie.Value
		}

		// Parse and validate the token
		tokenString := headerStr[len("Bearer "):]
		token, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.WriteError(w, errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		// Extract the user information from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.WriteError(w, errors.New("invalid token claims"), http.StatusUnauthorized)
			return
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			utils.WriteError(w, errors.New("invalid user id"), http.StatusUnauthorized)
			return
		}

		// Add the user information to the context
		ctx := context.WithValue(r.Context(), UserIDKey, userId)
		req := r.WithContext(ctx)

		// If the user is authenticated, call the next handler with the new context
		next.ServeHTTP(w, req)
	})
}
