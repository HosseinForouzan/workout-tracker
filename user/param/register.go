package param

type RegisterReqeust struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Name string `json:"name"`
}