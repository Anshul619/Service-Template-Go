package api

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
