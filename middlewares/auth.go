package middlewares

import (
	"context"
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/utils"
	"github.com/golang-jwt/jwt/v5"
)

type key string

const (
	UserIDKey key = "userId"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This should support web and mobile clients
		// Check if the user is authenticated

		// Check the Authorization header
		tokenStr := r.Header.Get("Authorization")

		// If the token is not in the Authorization header, check the cookies
		if tokenStr == "" {
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			tokenStr = cookie.Value
		}

		// Parse and validate the token
		// Replace "yourSigningKey" with your actual signing key
		token, err := utils.VerifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract the user information from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userId, ok := claims["userId"].(string)
		if !ok {
			http.Error(w, "Invalid user id", http.StatusUnauthorized)
			return
		}

		// Add the user information to the context
		ctx := context.WithValue(r.Context(), UserIDKey, userId)
		req := r.WithContext(ctx)

		// If the user is authenticated, call the next handler with the new context
		next.ServeHTTP(w, req)
	})
}
