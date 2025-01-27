package oauth

import (
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

func (o *Oauth2google) GetAuthURL() string {
	return o.config.AuthCodeURL("random")
}

//Todo: Implement the Authenticate method

func (o *Oauth2google) Authenticate(code string) (*domain.User, error) {
	// token, err := o.config.Exchange(context.Background(), code)
	// if err != nil {
	// 	return "", err
	// }
	// client := o.config.Client(context.Background(), token)
	// resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	// if err != nil {
	// 	return "", err
	// }
	// defer resp.Body.Close()
	// response, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }
	// return string(response), nil
	return &domain.User{}, nil
}
