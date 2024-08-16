package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id        string 
	Role      string 
	FirstName string 
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

}


type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

