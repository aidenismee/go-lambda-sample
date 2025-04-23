package contract

type GetUserRequest struct {
	ID int `param:"id" json:"id"`
}

type GetUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
