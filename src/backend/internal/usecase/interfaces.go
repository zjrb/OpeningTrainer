package usecase

type (
	AuthRepo interface {
	}

	Authenticate interface {
		AuthGoogle(code string) (string, error)
		GetGoogleAuthUrl() string
	}

	AuthWebApi interface {
		GetGoogleAuthUrl() string
		Auth(code string) (string, error)
	}
)
