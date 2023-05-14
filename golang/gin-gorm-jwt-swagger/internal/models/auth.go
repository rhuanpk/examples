package models

// Auth is the struct to manipule authentication operations.
type Auth struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password" binding:"required,min=8"`
}

// Token is the to representate model of token request.
type Token struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}
