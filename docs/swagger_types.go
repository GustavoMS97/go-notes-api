package docs

type ErrorResponse struct {
	Error string `json:"error" example:"Something went wrong"`
}

// login
type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"areallylongjwttoken"`
	RefreshToken string `json:"refresh_token" example:"areallylongjwttoken"`
}
