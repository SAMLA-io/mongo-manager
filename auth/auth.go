package auth

import (
	"context"
	"log"
	"mongo-manager/clerk"
	"net/http"
	"time"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

// OrganizationIDKey is the context key for storing organization ID
type OrganizationIDKey struct{}
type UserIDKey struct{}

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// VerifyingMiddleware is the general middleware that verifies the passed JWT Token from clerk and extracts the user ID and organization ID to pass it to the next handler
func VerifyingMiddleware(next http.Handler) http.Handler {
	return clerkhttp.RequireHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[AUTH] Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		startTime := time.Now()

		// Log authorization header presence (without revealing the token)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Printf("[AUTH] ERROR: Missing authorization header for %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("[AUTH] Authorization header present for %s %s", r.Method, r.URL.Path)

		userID, err := ExtractUserIDFromAuthHeader(r)
		if err != nil {
			log.Printf("[AUTH] ERROR: Failed to extract user ID for %s %s: %v", r.Method, r.URL.Path, err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("[AUTH] Successfully extracted user ID: %s for %s %s", userID, r.Method, r.URL.Path)

		organizationID, err := clerk.GetUserOrganizationId(userID)
		if err != nil {
			log.Printf("[AUTH] ERROR: Failed to get organization ID for user %s on %s %s: %v", userID, r.Method, r.URL.Path, err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("[AUTH] Successfully retrieved organization ID: %s for user %s on %s %s", organizationID, userID, r.Method, r.URL.Path)

		// Add organization ID and user ID to request context
		ctx := context.WithValue(r.Context(), OrganizationIDKey{}, organizationID)
		ctx = context.WithValue(ctx, UserIDKey{}, userID)
		r = r.WithContext(ctx)

		// Wrap the response writer to capture the status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("[AUTH] Response: %s %s -> STATUS: %d completed in %v (User: %s, Org: %s)", r.Method, r.URL.Path, rw.statusCode, time.Since(startTime), userID, organizationID)
	}))
}

// USED ONLY FOR TESTING
func TestingMiddleware(next http.Handler) http.Handler {
	return (http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[AUTH] Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		startTime := time.Now()

		organizationID := "org_2v0ixOVzW5VsqlFr1ZFJbc7GTay"
		log.Printf("[AUTH] Successfully retrieved organization ID: %s on %s %s", organizationID, r.Method, r.URL.Path)

		// Add organization ID to request context
		ctx := context.WithValue(r.Context(), OrganizationIDKey{}, organizationID)
		r = r.WithContext(ctx)

		// Wrap the response writer to capture the status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("[AUTH] Response: %s %s -> STATUS: %d completed in %v (Org: %s)", r.Method, r.URL.Path, rw.statusCode, time.Since(startTime), organizationID)
	}))
}
