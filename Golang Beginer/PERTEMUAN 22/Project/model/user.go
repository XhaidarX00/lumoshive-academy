package model

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// curl -X POST http://localhost:8080/api/auth/register -H "Content-Type: application/json" -d '{"username":"dev3", "password":"password789", "email":"dev3@example.com", "role":"dev"}'
