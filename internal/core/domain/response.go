package domain

type GoogleResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
}

type OAuthResponse struct {
	Email          string
	OAuthID        string
	OAuthProvider  string
	ProfilePicture string
	Name           string
}
