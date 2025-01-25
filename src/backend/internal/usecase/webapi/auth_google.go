package webapi

import (
	"backend/pkg/google"

	"context"
	"io"

	"golang.org/x/oauth2"
)

type AuthWebApi struct {
	*oauth2.Config
}

func New(g *google.Google) *AuthWebApi {
	return &AuthWebApi{g.Config}
}
func (a *AuthWebApi) GetGoogleAuthUrl() string {
	return a.Config.AuthCodeURL("random")
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
