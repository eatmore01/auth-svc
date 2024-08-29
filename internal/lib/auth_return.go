package lib

type LoginResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email     string `json:"email"`
	User_name string `json:"user_name"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
	Status int    `json:"status"`
	UserId string `json:"user_id"`
}
