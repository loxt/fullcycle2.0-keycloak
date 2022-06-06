package main

import (
	"context"
	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var (
	clientID     = "myclient"
	clientSecret = "HvoD5o6fPEoBPc4hkMxskHQgltdRtwo6"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/realms/myrealm")
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "foobar" // Don't do this in production.

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
