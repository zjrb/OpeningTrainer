package domain

type User struct {
	ID             int
	Name           string
	Email          string
	OAuthProvider  string
	OAuthID        string
	ProfilePicture string
	CreatedAt      string
	UpdatedAt      string
}
type contextKey string

const EmailContextKey = contextKey("email")

type Opening struct {
	OpeningName string
	ECO         string
	PGN         string
	UCI         string
	FEN         string
}

type GameSesion struct {
	Opening  []int  `redis:"opening"`
	White    bool   `redis:"white"`
	MoveNum  int    `redis:"moveNum"`
	LastMove string `redis:"lastMove"`
}
