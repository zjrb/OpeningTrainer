package usecase

type AuthUseCase struct {
	repo   AuthRepo
	webApi AuthWebApi
}

func New(r AuthRepo, w AuthWebApi) *AuthUseCase {
	return &AuthUseCase{
		repo:   r,
		webApi: w,
	}
}

func (a *AuthUseCase) GetGoogleAuthUrl() string {
	return a.webApi.GetGoogleAuthUrl()
}

func (a *AuthUseCase) AuthGoogle(code string) (string, error) {
	return a.webApi.Auth(code)
}
