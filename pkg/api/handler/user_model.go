package handler

type UserResponse struct{
	Data []User `json:"data"`
}

type User struct {
	UserId       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"` 
	ExpireDate   string    `json:"expire_date"` 
}
