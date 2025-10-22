package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	clerkjwt "github.com/clerk/clerk-sdk-go/v2/jwt"
)

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}

// GetOrganizationID retrieves the organization ID from the request context
func GetOrganizationID(r *http.Request) (string, bool) {
	organizationID, ok := r.Context().Value(OrganizationIDKey{}).(string)
	return organizationID, ok
}

func GetUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(UserIDKey{}).(string)
	return userID, ok
}

// extractUserIDFromAuthHeader extracts the user ID from the Authorization header
func ExtractUserIDFromAuthHeader(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing authorization header")
	}

	// Check if it's a Bearer token
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header format")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Verify the JWT token and extract claims
	claims, err := clerkjwt.Verify(context.Background(), &clerkjwt.VerifyParams{
		Token: token,
	})
	if err != nil {
		return "", fmt.Errorf("failed to verify token: %v", err)
	}

	// Extract user ID from the subject claim
	userID := claims.RegisteredClaims.Subject
	if userID == "" {
		return "", fmt.Errorf("no user ID found in token")
	}

	return userID, nil
}
