package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"golang.org/x/oauth2"
)

type Oauth2google struct {
	config *oauth2.Config
}

func NewOauth2google(config *oauth2.Config) *Oauth2google {
	return &Oauth2google{
		config: config,
	}
}

func (o *Oauth2google) GetAuthURL(state string) string {
	return o.config.AuthCodeURL(state)
}

func (o *Oauth2google) Authenticate(code string) (*domain.OAuthResponse, error) {
	token, err := o.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	client := o.config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var googleResponse domain.GoogleResponse
	response, err := io.ReadAll(resp.Body)
	fmt.Println("Google Response: " + string(response))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &googleResponse)
	if err != nil {
		return nil, err
	}
	return &domain.OAuthResponse{OAuthID: googleResponse.ID, Email: googleResponse.Email,
		ProfilePicture: googleResponse.Picture, OAuthProvider: "google", Name: googleResponse.Name}, nil
}

func (o *Oauth2google) GenerateStateOauthCookie() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}
