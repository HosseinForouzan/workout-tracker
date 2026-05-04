package param

type ProfileRequest struct {
	UserID uint 
}

type ProfileResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
}