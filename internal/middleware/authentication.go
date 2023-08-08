package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	authentication "github.com/deanrtaylor1/backend-go/internal/Auth"
	db "github.com/deanrtaylor1/backend-go/internal/db/sqlc"
)

func AuthMiddleware(authenticator authentication.Authenticator, store db.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawToken, err := extractToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			verifier := authenticator.Provider.Verifier(&oidc.Config{ClientID: authenticator.Config.ClientID})

			idToken, err := verifier.Verify(r.Context(), rawToken)
			if err != nil {
				fmt.Println("no idToken found")
				fmt.Println(err)
				next.ServeHTTP(w, r)
				return
			}
			fmt.Println(idToken)

			var claims struct {
				Sub string `json:"sub"`
			}
			if err := idToken.Claims(&claims); err != nil {
				next.ServeHTTP(w, r)
				return
			}
			fmt.Println(claims)
			fmt.Println(idToken)

			next.ServeHTTP(w, r)
		})
	}
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no Authorization header provided")

	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("malformed Authorization header")
	}

	return parts[1], nil
}
