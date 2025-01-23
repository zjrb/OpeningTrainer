package webapi

import (
	"backend/pkg/google"

	"context"
	"io"

	"golang.org/x/oauth2"
)

type AuthWebApi struct {
	*google.Google
}

func New(clientId string, clientSecret string, redirectUrl string) *AuthWebApi {
	return &AuthWebApi{google.New(google.ClientID(clientId), google.ClientSecret(clientSecret), google.RedirectURL(redirectUrl))}
}
func (a *AuthWebApi) GetGoogleAuthUrl() string {
	return a.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}
func (a *AuthWebApi) Auth(code string) (string, error) {
	token, err := a.Config.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}
	client := a.Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
