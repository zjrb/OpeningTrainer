package google

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	_defaultScopes = "email"
)

type Google struct {
	clientId     string
	clientSecret string
	redirectUrl  string
	scopes       []string
	Config       *oauth2.Config
}

func New(opts ...Option) *Google {
	g := &Google{
		Config: &oauth2.Config{
			Scopes: []string{_defaultScopes},
		},
	}
	for _, opt := range opts {
		opt(g)
	}
	g.Config.Endpoint = google.Endpoint
	fmt.Printf("Redirect URL: %s\n", g.Config.RedirectURL)
	return g
}
