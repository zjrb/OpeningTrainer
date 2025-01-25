package google

type Option func(*Google)

func Scopes(scopes ...string) Option {
	return func(g *Google) {
		g.Config.Scopes = scopes
	}
}

func RedirectURL(url string) Option {
	return func(g *Google) {
		g.Config.RedirectURL = url
	}
}
func ClientID(id string) Option {
	return func(g *Google) {
		g.Config.ClientID = id
	}
}

func ClientSecret(secret string) Option {
	return func(g *Google) {
		g.Config.ClientSecret = secret
	}
}
