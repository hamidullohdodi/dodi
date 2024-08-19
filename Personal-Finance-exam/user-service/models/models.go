package models

type User struct {
	Id           string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Phone        string `json:"phone_number"`
	Image        string `json:"image"`
	Role         string `json:"role"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
type Register struct {
	FirstName     string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Phone        string `json:"phone_number"`
	Image        string `json:"image"`
	Role         string `json:"role"`
}
type RegisterResponse struct {
	Id           string `json:"id"`
	FirstName     string `json:"user_name"`
	LastName     string `json:"full_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Phone        string `json:"phone"`
	Image        string `json:"image"`
	Role         string `json:"role"`
	CreatedAt    string `json:"created_at"`
}
type LoginRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
type LoginResponse struct {
	Id           string `json:"id"`
	FirstName     string `json:"user_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}
type Token struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	ExpiredTime  float64 `json:"expired_time"`
}
type SendEmail struct {
	Email              string `json:"email"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type SaveToken struct {
	UserId  string
	Token   string
	Revoked bool
}
