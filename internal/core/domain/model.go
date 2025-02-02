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
	MoveArray   []string
}

type GameSession struct {
	Opening  []string `redis:"opening" json:"opening"`
	White    bool     `redis:"white" json:"white"`
	MoveNum  int      `redis:"moveNum" json:"moveNum"`
	LastMove string   `redis:"lastMove" json:"lastMove"`
	Message  string   `json:"message"`
}

type GameSessionDB struct {
	ID         int
	OpeningID  int
	UserID     int
	WrongMoves int
	Accuracy   float32
}
