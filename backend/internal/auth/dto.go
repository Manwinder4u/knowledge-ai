package auth

// Data Transfer Objects

// User Register request object
type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// User Login request object
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User Response object
type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
